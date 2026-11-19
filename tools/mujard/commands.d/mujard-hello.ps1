# A custom command: any executable named mujard-<name> becomes "mujard <name>".
# This one is the example from README.md - "mujard hello" runs it.
param([Parameter(ValueFromRemainingArguments = $true)][string[]]$Rest)
Write-Output "hello from a mujard plugin ($($Rest -join ' '))"
Write-Output "api: $env:MUJARD_API  root: $env:MUJARD_ROOT"
