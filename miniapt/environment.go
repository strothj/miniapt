package miniapt

import (
	"os"
	"path/filepath"

	"golang.org/x/crypto/openpgp"

	"github.com/strothj/go-debrepo/debrepo"
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
	DataDir          string
	ConfigDir        string
	LoadRepositories func() debrepo.RepositoryList
	SaveRepositories func(debrepo.RepositoryList) error
	SaveKey          func(openpgp.EntityList) error
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
	env.LoadRepositories = func() debrepo.RepositoryList {
		return LoadRepository(filepath.Join(env.ConfigDir, "sources.list"))
	}
	env.SaveRepositories = func(repoList debrepo.RepositoryList) error {
		err := os.MkdirAll(env.ConfigDir, 0755)
		if err != nil {
			return err
		}
		return SaveRepositories(filepath.Join(env.ConfigDir, "sources.list"), repoList)
	}
	env.SaveKey = func(entityList openpgp.EntityList) error {
		keysPath := filepath.Join(env.ConfigDir, "keys")
		err := os.MkdirAll(keysPath, 0755)
		if err != nil {
			return err
		}
		return SaveKey(keysPath, entityList)
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
