module Main where
import Control.Monad (liftM, when)
import Sound.ALSA.Mixer
import System.Environment (getArgs)

mixer = "default"
control = "Master"

volumeHelper = withMixer mixer $ \mixer -> do
  Just control <- getControlByName mixer control
  let Just playbackVolume = playback $ volume control
  (min, max) <- getRange playbackVolume
  return (playbackVolume, min, max)

getVolume = do
  (playbackVolume, min, max) <- volumeHelper
  Just vol <- getChannel FrontLeft $ value playbackVolume
  return vol

changeVolumeBy i = do
  (playbackVolume, min, max) <- volumeHelper
  Just vol <- getChannel FrontLeft $ value playbackVolume
  when ((i > 0 && vol < max) || (i < 0 && vol > min))
    $ setChannel FrontLeft (value playbackVolume) $ vol + i

setVolumeTo i = do
  (playbackVolume, min, max) <- volumeHelper
  when (i >= min || i <= max)
    $ setChannel FrontLeft (value playbackVolume) i

muteHelper = withMixer mixer $ \mixer -> do
  Just control <- getControlByName mixer control
  let Just playbackSwitch = playback $ switch control
  Just sw <- getChannel FrontLeft playbackSwitch
  return (playbackSwitch, sw)

getMute = liftM snd muteHelper

toggleMute = muteHelper >>= \(playbackSwitch, sw) -> setChannel FrontLeft playbackSwitch $ not sw

main = do
  args <- getArgs
  case args of
    [] -> return ()
    "toggle":_  -> toggleMute
    ('+':num):_ -> changeVolumeBy $ read num
    ('-':num):_ -> changeVolumeBy $ negate $ read num
    num:_       -> setVolumeTo $ read num
  vol <- getVolume
  on <- getMute
  putStrLn (show vol ++ ", " ++ if on then "on" else "off")

