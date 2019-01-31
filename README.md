# volume

A tool for easily querying and manipulating the ALSA master volume from the
command line.

Mute, un-mute, toggle mute using
``` sh
volume mute
volume unmute
volume toggle
```

Increment by 4, decrement by 3 or set to 57 using
``` sh
volume +4
volume -3
volume 57
```

For use with xmobar, showing muted/not muted using font-awesome, use

``` sh
volume -x
```
