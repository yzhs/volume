#!/usr/bin/python3

import sys
from subprocess import run, PIPE

mixer = "Master"


def amixer(command, mixer=mixer, silent=False):
    """Execute a command using amixer."""
    cmd = None
    if command == "get":
        cmd = ["amixer", "get", mixer]
    else:
        cmd = ["amixer", "set", mixer, command]
        if silent:
            cmd.insert(1, "-q")
    output = run(cmd, stdout=PIPE).stdout.decode("utf-8")

    # Parse output
    pattern = "Mono: Playback "
    i = output.find(pattern) + len(pattern)
    output = output[i:]
    i = output.find(" ")
    volume = output[:i]
    output = output[i:]
    pattern = "dB] ["
    i = output.find(pattern) + len(pattern)
    mute = output[i:]
    mute = mute[:mute.find("]")]
    return (volume, mute)


def handle_one_arg(arg):
    """Handle any of the arguments related to setting or getting the
    volume."""
    result = None
    if arg == "toggle" or arg == "unmute":
        result = amixer(arg)
        # Work around a bug (?) in PulseAudio
        amixer("unmute", "Headphone", silent=True)
        amixer("unmute", "Speaker", silent=True)
    elif arg == "mute" or arg[0] in "1234567890":
        result = amixer(arg)
    elif arg[0] in "+-":
        result = amixer(arg[1:] + arg[0])
    return result


def print_result(result, use_gui=False):
    if result is None:
        return

    mute = result[1]
    if use_gui and mute == "on":
        mute = "<fn=1></fn>"
    elif use_gui and mute == "off":
        mute = "<fn=1></fn> "
    if use_gui:
        print("<fc=#cccc00>{}{}</fc>".format(mute, result[0]))
    else:
        print(result[0], mute)


result = None
args = sys.argv[1:]
if len(args) == 0:
    result = amixer("get")
elif len(args) == 1:
    result = handle_one_arg(args[0])

print_result(result, gui)
