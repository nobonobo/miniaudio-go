//go:build windows

package miniaudio

import (
	"embed"
	"fmt"
	"runtime"

	"golang.org/x/sys/windows"
)

//go:embed build/windows/*
var buildEmbed embed.FS

func openLibrary() (uintptr, error) {
	name := "build/windows/libminiaudio-windows-" + runtime.GOARCH + ".so"

	dynamicLibrary, err := windows.LoadLibrary(name)
	if err != nil {
		return 0, fmt.Errorf("loading windows dll: %w", err)
	}

	return uintptr(dynamicLibrary), nil
}
