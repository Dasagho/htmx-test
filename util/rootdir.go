package util

import (
	"os"
	"path/filepath"
)

func FindGoMod() (string, error) {
	dir, err := os.Getwd()
	if err != nil {
		return "", err
	}

	for dir != "/" { // Cambia esto a `for dir != filepath.VolumeName(dir) + "\\"` en Windows.
		if _, err := os.Stat(filepath.Join(dir, "go.mod")); err == nil {
			return dir, nil
		}

		dir = filepath.Dir(dir)
	}

	return "", os.ErrNotExist
}
