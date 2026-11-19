#!/bin/sh
# A custom command: any executable named mujard-<name> becomes "mujard <name>".
# This one is the example from README.md - "mujard hello" runs it.
echo "hello from a mujard plugin ($*)"
echo "api: $MUJARD_API  root: $MUJARD_ROOT"
