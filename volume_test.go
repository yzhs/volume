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
