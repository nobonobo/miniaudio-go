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

- Fix deadlock when not setting any device formats (should use device defaults).
    - Issue is caused by the fact that we were using config to know channel.
    - Need to find a way to set dta callback afte init, so we can know channel count within the data callback.
- Read number of channels from miniaudio context.
- Generate intermittant C file after macros.
