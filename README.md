<p align="center">
  <img src="docs/logo.png" alt="suppress_success" width="200">
</p>

<h1 align="center">suppress_success</h1>

<p align="center"><em>Noise-free by design.</em></p>

<p align="center">A CLI that suppresses output on success. On failure, it prints the full output for debugging and returns the original exit code.</p>

[![Build](https://github.com/isidore/SuppressSuccess/actions/workflows/build.yml/badge.svg)](https://github.com/isidore/SuppressSuccess/actions/workflows/build.yml)

## Why 
Too often we pollute AI with reams of output from our builds.
While this is needed to debug failures, it is clutter when the build passes.
suppress_success is an easy way to get the best of both.

## Quick start

```bash
suppress_success echo "Hello"                    # prints "Success"
suppress_success --message "Hi" echo "Hello"     # prints "Hi"
suppress_success sh -c 'echo "Hello" && exit 1'  # prints "Hello", exits 1
```


## Installation

Download the latest binary from [GitHub Releases](https://github.com/isidore/suppress_success/releases).

```bash
# Mac Apple Silicon
curl -L https://github.com/isidore/suppress_success/releases/latest/download/suppress_success-mac-apple-silicon -o suppress_success
chmod +x suppress_success
mv suppress_success /usr/local/bin/
```

```bash
# Mac Intel
curl -L https://github.com/isidore/suppress_success/releases/latest/download/suppress_success-mac-intel -o suppress_success
chmod +x suppress_success
mv suppress_success /usr/local/bin/
```

```powershell
# Windows
curl -L https://github.com/isidore/suppress_success/releases/latest/download/suppress_success-windows.exe -o suppress_success.exe
```

## Usage

See the [User Guide](docs/guide.md) for full documentation.
