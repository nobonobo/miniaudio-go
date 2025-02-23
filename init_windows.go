//go:build windows

package miniaudio

import (
	"embed"
	"fmt"
	"io/fs"
	"path"
	"runtime"
	"strings"

	"golang.org/x/sys/windows"
)

//go:embed build/windows/*
var buildEmbed embed.FS

const rootDir = "build/windows"

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

		dynamicLibrary, err := windows.LoadLibrary(path.Join(rootDir, dirEntry.Name()))
		if err != nil {
			return 0, fmt.Errorf("loading windows dll: %w", err)
		}

		return uintptr(dynamicLibrary), nil
	}

	return 0, ErrLibraryNotFound
}
