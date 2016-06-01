package miniapt

import (
	"os"
	"path/filepath"

	"github.com/strothj/debrepo"
)

const (
	// DataDirFlagName is the command line flag for the application data
	// directory.
	DataDirFlagName = "datadir"
)

// StringFlagger provides the string value of a command line flag.
type StringFlagger interface {
	String(name string) string
}

// Environment provides application services to the command line interface.
type Environment struct {
	DataDir     string
	ConfigDir   string
	LoadSources func() debrepo.SourceList
	SaveSources func(debrepo.SourceList) error
}

// EnvironmentFromContext returns an Environment configured using values passed
// to the command line interface.
func EnvironmentFromContext(ctx StringFlagger) (*Environment, error) {
	env := new(Environment)
	dataDir, err := getDataDirPath(ctx)
	if err != nil {
		return nil, err
	}
	env.DataDir = dataDir
	env.ConfigDir = filepath.Join(dataDir, "config")
	env.LoadSources = func() debrepo.SourceList {
		return LoadSources(filepath.Join(env.ConfigDir, "sources.list"))
	}
	env.SaveSources = func(sourceList debrepo.SourceList) error {
		err := os.MkdirAll(env.ConfigDir, 0755)
		if err != nil {
			return err
		}
		err = SaveSources(filepath.Join(env.ConfigDir, "sources.list"), sourceList)
		if err != nil {
			return err
		}
		return nil
	}
	return env, nil
}

func getDataDirPath(ctx StringFlagger) (string, error) {
	if path := ctx.String(DataDirFlagName); len(path) > 0 {
		return filepath.Clean(path), nil
	}
	d, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Join(d, "miniapt"), nil
}
