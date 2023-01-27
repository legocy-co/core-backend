package helpers

import (
	"errors"
	"os"
)

func FileExists(fp string) bool {
	_, err := os.Stat(fp)
	return !errors.Is(err, os.ErrNotExist)
}
