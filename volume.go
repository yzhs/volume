package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const mixer = "Master"

func main() {
	var output string

	gui := false
	args := os.Args[1:]

	if len(args) == 0 || (len(args) == 1 && args[0] == "-x") {
		if len(args) > 0 {
			gui = true
		}
		output = call_amixer("get", mixer, false)
	} else if len(args) >= 1 {
		if len(args) > 1 && args[0] == "-x" {
			gui = true
			args = args[1:]
		}

		arg := args[0]

		if arg == "toggle" || arg == "unmute" {
			output = call_amixer(arg, mixer, false)
			// Work around a bug (?) in PulseAudio
			call_amixer("unmute", "Headphone", true)
			call_amixer("unmute", "Speaker", true)
		} else if arg == "mute" || strings.IndexByte("1234567890", arg[0]) != -1 {
			output = call_amixer(arg, mixer, false)
		} else if arg[0] == '-' || arg[0] == '+' {
			output = call_amixer(arg[1:]+arg[:1], mixer, false)
		}
	}

	vol, muted := parse_output(output)

	print_result(vol, muted, gui)
}

func call_amixer(command string, mixer string, silent bool) string {
	arg := ""
	if silent {
		arg = "-q"
	}

	var cmd string
	if command == "get" {
		cmd = "get"
	} else {
		cmd = "set"
	}

	tmp, err := exec.Command("amixer", cmd, mixer, command, arg).CombinedOutput()
	if err != nil {
		panic(err)
	}

	return string(tmp)
}

func parse_output(output string) (volume, mute string) {
	pattern := "Mono: Playback "
	i := strings.Index(output, pattern) + len(pattern)
	output = output[i:]
	i = strings.IndexRune(output, ' ')
	volume = output[:i]
	output = output[i:]

	pattern = "dB] ["
	i = strings.Index(output, pattern) + len(pattern)
	mute = output[i:]
	mute = mute[:strings.Index(mute, "]")]
	return volume, mute
}

func print_result(volume, mute string, use_gui bool) {
	if use_gui && mute == "on" {
		mute = "<fc=#00aa00><fn=1></fn></fc>"
	} else if use_gui && mute == "off" {
		mute = "<fc=red><fn=1></fn></fc> "
	}
	if use_gui {
		println(mute, volume)
	} else {
		fmt.Printf("%s %s\n", volume, mute)
	}
}
