# miniaudio-go [WIP]

Current C library version: v0.11.22

## Description

This library provides a wrapper around the cross-platform [miniaudio](https://github.com/mackron/miniaudio) C library.
It uses [purego](https://github.com/ebitengine/purego) to load static builds of the miniaudio library.
The miniaudio library is cross-compiled using [zig](https://ziglang.org) (v0.13.0) as C compiler.
The static binary of the target platform is embedded into the Go binary when importing this package.
This allows for static Go binaries which can be compiled with `CGO_ENABLED=0`.

Special effort was made to translate miniaudio's C API to a Go API which should feel idiomatic to Go developers, while remaining familiar for miniaudio users.

The current scope of the project is not to provide one-to-one feature parity with miniaudio, but to provide a cross-platform audio solution for Go developers without a dependency on CGO.
Initial focus is on Linux & Windows operating systems with `amd64` architecture.

## TODO

- Find way to set data callback after device init, so format & channels can be read from device info.
- Generate intermittant C file after macros.
