package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const mixer = "Master"

func main() {
	arg, unmute, gui := ParseArguments(os.Args[1:])

	output := CallAmixer(arg, mixer)

	if unmute {
		unmuteIndividualOutputs()
	}

	vol, muted := ParseOutput(output)

	PrintResult(vol, muted, gui)
}

// Parse command line arguments
func ParseArguments(args []string) (arg string, unmute, gui bool) {
	gui = len(args) > 0 && args[0] == "-x"
	if gui {
		args = args[1:]
	}

	arg = "get"
	if len(args) > 0 {
		arg = args[0]
	}

	switch {
	case arg == "toggle", arg == "unmute":
		unmute = true
	case arg[0] == '-' || arg[0] == '+':
		arg = arg[1:] + arg[:1]
	case arg == "get" || arg == "mute" || '0' <= arg[0] && arg[0] <= '9':
	default:
		panic("Unknown argument")
	}

	return arg, unmute, gui
}

// Execute `amixer (get|set) <mixer> ...`,
func CallAmixer(command, mixer string) string {
	return callAmixerHelper(command, mixer, false)
}

func callAmixerQuiet(command, mixer string) {
	callAmixerHelper(command, mixer, true)
}

func callAmixerHelper(command, mixer string, silent bool) string {
	arg := ""
	if silent {
		arg = "-q"
	}

	cmd := command
	if cmd != "get" {
		cmd = "set"
	}

	tmp, err := exec.Command("amixer", cmd, mixer, command, arg).CombinedOutput()
	if err != nil {
		panic(err)
	}

	return string(tmp)
}

func ParseOutput(output string) (volume, mute string) {
	pattern := "Mono: Playback "
	i := strings.Index(output, pattern)
	if i == -1 {
		panic(output)
	}
	output = output[i+len(pattern):]
	i = strings.IndexRune(output, ' ')
	volume = output[:i]
	output = output[i:]

	pattern = "dB] ["
	i = strings.Index(output, pattern) + len(pattern)
	mute = output[i:]
	mute = mute[:strings.Index(mute, "]")]
	return volume, mute
}

func PrintResult(volume, mute string, useGui bool) {
	if useGui && mute == "on" {
		mute = "<fc=#00aa00><fn=1></fn></fc>"
	} else if useGui && mute == "off" {
		mute = "<fc=red><fn=1></fn></fc> "
	}
	if useGui {
		println(mute, volume)
	} else {
		fmt.Printf("%s %s\n", volume, mute)
	}
}

func unmuteIndividualOutputs() {
	// Work around a bug (?) in PulseAudio
	callAmixerQuiet("unmute", "Headphone")
	callAmixerQuiet("unmute", "Speaker")
}
