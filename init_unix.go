//go:build linux

package miniaudio

import (
	"fmt"

	"github.com/ebitengine/purego"
)

func openLibrary(name string) (uintptr, error) {
	dynamicLibrary, err := purego.Dlopen(name, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, fmt.Errorf("loading unix library: %w", err)
	}

	return dynamicLibrary, nil
}
