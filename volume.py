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
    output = run(cmd)


def handle_one_arg(arg):
    """Handle any of the arguments related to setting or getting the
    volume."""
    if arg == "toggle" or arg == "unmute":
        result = amixer(arg)
        # Work around a bug (?) in PulseAudio
        amixer("unmute", "Headphone", silent=True)
        amixer("unmute", "Speaker", silent=True)
    elif arg == "mute" or arg[0] in "1234567890":
        amixer(arg)
    elif arg[0] in "+-":
        amixer(arg[1:] + arg[0])


def print_result(result):
    if result is None:
        return

    print(result[0]+",", result[1])


result = None
args = sys.argv[1:]
if len(args) == 0:
    amixer("get")
elif len(args) == 1:
    handle_one_arg(args[0])

