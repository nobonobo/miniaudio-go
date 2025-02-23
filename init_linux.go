//go:build linux

package miniaudio

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"runtime"
	"strings"

	"github.com/ebitengine/purego"
)

//go:embed build/linux/*
var buildEmbed embed.FS

const rootDir = "build/linux"

func openLibrary() (uintptr, error) {
	dirEntries, err := fs.ReadDir(buildEmbed, rootDir)
	if err != nil {
		return 0, fmt.Errorf("reading dir: %w", err)
	}

	for _, dirEntry := range dirEntries {
		if dirEntry.IsDir() {
			continue
		}

		if !strings.Contains(dirEntry.Name(), runtime.GOARCH) {
			continue
		}

		dynamicLibrary, err := purego.Dlopen(path.Join(rootDir, dirEntry.Name()), purego.RTLD_NOW|purego.RTLD_GLOBAL)
		if err != nil {
			return 0, fmt.Errorf("loading unix library: %w", err)
		}

		return dynamicLibrary, nil
	}

	return 0, ErrLibraryNotFound
}
