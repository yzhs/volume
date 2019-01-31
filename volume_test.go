package main

import "testing"

func TestParseArgumentsTrivial(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{})

	if arg != "get" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsTrivialGui(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"-x"})

	if arg != "get" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if !gui {
		t.Fail()
	}
}

func TestParseArgumentsIncrement(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"+42"})

	if arg != "42+" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsSet(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"42"})

	if arg != "42" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsDecrement(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"-1234"})

	if arg != "1234-" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsGuiMute(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"-x", "mute"})

	if arg != "mute" {
		t.Fail()
	}

	if unmute {
		t.Fail()
	}

	if !gui {
		t.Fail()
	}
}

func TestParseArgumentsUnmute(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"unmute"})

	if arg != "unmute" {
		t.Fail()
	}

	if !unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsToggle(t *testing.T) {
	arg, unmute, gui := parseArguments([]string{"toggle"})

	if arg != "toggle" {
		t.Fail()
	}

	if !unmute {
		t.Fail()
	}

	if gui {
		t.Fail()
	}
}

func TestParseArgumentsFailure(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	parseArguments([]string{"foo"})
}

func TestParseOutputUnmutedVolume(t *testing.T) {
	output := "Simple mixer control 'Master',0\n  Capabilities: pvolume pvolume-joined pswitch pswitch-joined\n  Playback channels: Mono\n  Limits: Playback 0 - 74\n  Mono: Playback 58 [78%] [-16.00dB] [on]"

	volume, muted := parseOutput(output)

	if volume != "58" {
		t.Fail()
	}

	if muted != "on" {
		t.Fail()
	}
}

func TestParseOutputmuteddVolume(t *testing.T) {
	output := "Simple mixer control 'Master',0\n  Capabilities: pvolume pvolume-joined pswitch pswitch-joined\n  Playback channels: Mono\n  Limits: Playback 0 - 74\n  Mono: Playback 58 [78%] [-16.00dB] [off]"

	volume, muted := parseOutput(output)

	if volume != "58" {
		t.Fail()
	}

	if muted != "off" {
		t.Fail()
	}
}

func TestParseOutputStereo(t *testing.T) {
	output := "Simple mixer control 'Headphone',0\n  Capabilities: pvolume pswitch\n  Playback channels: Front Left - Front Right\n  Limits: Playback 0 - 74\n  Mono:\n  Front Left: Playback 72 [100%] [0.00dB] [on]\n  Front Right: Playback 73 [100%] [0.00dB] [off]"

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		}
	}()

	parseOutput(output)
}
