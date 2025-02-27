# [WIP] miniaudio

Current C library version: v0.11.22

## Description

This library provides a wrapper around the cross-platform [miniaudio](https://github.com/mackron/miniaudio) C library.
It uses [purego](https://github.com/ebitengine/purego) to load static builds of the miniaudio library.
The miniaudio library is cross-compiled using [zig](https://ziglang.org) (v0.13.0) as C compiler.
The static binaries are embedded into your Go binary when importing this package.
This allows for static Go binaries which can be compiled with `CGO_ENABLED=0`.

Special effort was made to translate miniaudio's C API to a Go API which should feel idiomatic to Go developers, while remaining familiar for miniaudio users.

The current scope of the project is not to provide one-to-one feature parity with miniaudio, but to provide a cross-platform audio solution for Go developers.
Initial focus is on Linux & Windows operating systems with `amd64` architectures as I have access to those platforms. 

## Help wanted!

Does your machine run MacOS? I am looking for a static miniaudio binary for `arm64` and `amd64`.
I am unable to cross-compile a MacOS build using Zig due to a `CoreAudio` dependency.
If someone could build these binaries for me I can add MacOS support to this project!

## TODO

- Fix deadlock when not setting any device formats (should use device defaults).
- Read number of channels from miniaudio context.
- Generate intermittant C file after macros.
