//go:build linux

package miniaudio

import (
	"embed"
	"fmt"
	"runtime"

	"github.com/ebitengine/purego"
)

//go:embed build/linux/*
var buildEmbed embed.FS

func openLibrary() (uintptr, error) {
	name := "build/linux/libminiaudio-linux-" + runtime.GOARCH + ".so"

	dynamicLibrary, err := purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, fmt.Errorf("loading unix library: %w", err)
	}

	return dynamicLibrary, nil
}
