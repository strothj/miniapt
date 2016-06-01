package miniapt

import (
	"os"
	"path/filepath"
	"testing"
)

type getDataDirTestFlagger struct {
	path string
}

func (g *getDataDirTestFlagger) String(name string) string {
	if name != DataDirFlagName {
		return ""
	}
	return g.path
}

func TestGetDataDirPath_DefaultsToPWD(t *testing.T) {
	pwd, err := os.Getwd()
	if err != nil {
		t.Fatal(err)
	}
	dir, err := GetDataDirPath(&getDataDirTestFlagger{})
	if err != nil {
		t.Fatal(err)
	}
	if expected, actual := filepath.Join(pwd, "miniapt"), dir; expected != actual {
		t.Fatalf("expected=%v actual=%v", expected, actual)
	}
}
