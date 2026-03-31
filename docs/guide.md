# User Guide

`suppress_success` is a CLI that runs a command and suppresses its output on success. On failure it prints the full output and exits with the same exit code as the command. On success it prints a short message ("Success" by default) and exits 0.

## Example Usage

```bash
suppress_success echo "Hello"                     # prints "Success"
suppress_success --message "Hi" echo "Hello"      # prints "Hi"
suppress_success sh -c 'echo "Hello" && exit 1'   # prints "Hello", exits 1
```

## Options

| Flag | Description |
|---|---|
| `--message <text>` | Message to print on success (default: `Success`) |

## Installation

Download the binary for your platform, make it executable, and put it on your PATH:

```bash
chmod +x suppress_success
mv suppress_success /usr/local/bin/
```
