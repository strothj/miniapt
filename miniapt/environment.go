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
	ctx       StringFlagger
	DataDir   string
	ConfigDir string
}

// EnvironmentFromContext returns an Environment configured using values passed
// to the command line interface.
func EnvironmentFromContext(ctx StringFlagger) (*Environment, error) {
	dataDir, err := getDataDirPath(ctx)
	if err != nil {
		return nil, err
	}
	return &Environment{
		ctx:       ctx,
		DataDir:   dataDir,
		ConfigDir: filepath.Join(dataDir, "config"),
	}, nil
}

// LoadRepositories reads the list of repositories from the data directory.
func (e *Environment) LoadRepositories() debrepo.RepositoryList {
	return LoadRepository(filepath.Join(e.ConfigDir, "sources.list"))
}

// SaveRepositories saves the list of repositories to the data directory.
func (e *Environment) SaveRepositories(repoList debrepo.RepositoryList) error {
	err := os.MkdirAll(e.ConfigDir, 0755)
	if err != nil {
		return err
	}
	return SaveRepositories(filepath.Join(e.ConfigDir, "sources.list"), repoList)
}

// SaveKey saves an OpenPGP key to the data directory.
func (e *Environment) SaveKey(entityList openpgp.EntityList) error {
	keysPath := filepath.Join(e.ConfigDir, "keys")
	err := os.MkdirAll(keysPath, 0755)
	if err != nil {
		return err
	}
	return SaveKey(keysPath, entityList)
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
