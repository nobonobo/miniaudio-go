//go:build windows

package miniaudio

import (
	"fmt"

	"golang.org/x/sys/windows"
)

func openLibrary(name string) (uintptr, error) {
	dynamicLibrary, err := windows.LoadLibrary(name)
	if err != nil {
		return 0, fmt.Errorf("loading windows dll: %w", err)
	}

	return uintptr(dynamicLibrary), nil
}
