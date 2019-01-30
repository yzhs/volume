package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

const mixer = "Master"

func main() {
	var vol, muted string

	gui := false
	args := os.Args[1:]

	if len(args) == 0 || (len(args) == 1 && args[0] == "-x") {
		if len(args) > 0 {
			gui = true
		}
		vol, muted = call_amixer("get", mixer, false, true)
	} else if len(args) >= 1 {
		if len(args) > 1 && args[0] == "-x" {
			gui = true
			args = args[1:]
		}

		arg := args[0]

		if arg == "toggle" || arg == "unmute" {
			vol, muted = call_amixer(arg, mixer, false, true)
			// Work around a bug (?) in PulseAudio
			call_amixer("unmute", "Headphone", true, false)
			call_amixer("unmute", "Speaker", true, false)
		} else if arg == "mute" || strings.IndexByte("1234567890", arg[0]) != -1 {
			vol, muted = call_amixer(arg, mixer, false, true)
		} else if arg[0] == '-' || arg[0] == '+' {
			vol, muted = call_amixer(arg[1:]+arg[:1], mixer, false, true)
		}
	}

	print_result(vol, muted, gui)
}

// Execute a command using amixer.
func call_amixer(command string, mixer string, silent, need_output bool) (volume, mute string) {
	arg := ""
	if silent {
		arg = "-q"
	}
	cmd := "set"
	if command == "get" {
		cmd = command
	}
	fmt.Println("amixer", cmd, mixer, command, arg)
	tmp, err := exec.Command("amixer", cmd, mixer, command, arg).CombinedOutput()
	if err != nil {
		panic(err)
	}

	if !need_output {
		return "", ""
	}

	output := string(tmp)

	// Parse output
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
