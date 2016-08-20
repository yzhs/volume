#!/usr/bin/python

import sys
from subprocess import call

mixer = "Master"

def call_amixer(cmd, args, silent=False):
    if silent:
        call(["amixer", "-q", cmd] + list(args))
    else:
        call(["amixer", cmd] + list(args))

args = sys.argv[1:]
if len(args) == 0:
    call_amixer("get", ["Master"])
elif len(args) == 1:
    if args[0] == "toggle":
        call_amixer("set", ["Master", "toggle"])
        call_amixer("set", ["Headphone", "unmute"], silent=True)
        call_amixer("set", ["Speaker", "unmute"], silent=True)
    elif args[0][0] in "1234567890":
        call_amixer("set", ["Master", args[0]])
    elif args[0][0] in "+-":
        call_amixer("set", ["Master", args[0][1:] + args[0][0]])

