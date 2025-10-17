//go:build darwin

package miniaudio

import (
	"fmt"

	"github.com/ebitengine/purego"
)

func openLibrary(fileName string) (uintptr, error) {
	dynamicLibrary, err := purego.Dlopen(fileName, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return 0, fmt.Errorf("loading darwin library: %w", err)
	}

	return dynamicLibrary, nil
}
