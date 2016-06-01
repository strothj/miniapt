package miniapt

import (
	"os"
	"path/filepath"
)

// StringFlagger provides the string value of a command line flag.
type StringFlagger interface {
	String(name string) string
}

// GetDataDirPath returns the application data directory. Returns $PWD/miniapt
// if the environment variable and command line flag are not set. Returns an
// error if the current working directory could not be read.
func GetDataDirPath(ctx StringFlagger) (string, error) {
	if path := ctx.String(DataDirFlagName); len(path) > 0 {
		return filepath.Clean(path), nil
	}
	d, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, "miniapt"), nil
}
