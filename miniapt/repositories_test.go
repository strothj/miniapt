package miniapt

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"testing"

	"github.com/strothj/go-debrepo/debrepo"
)

var repositoryTest = []string{
	"deb http://ftp.debian.org/debian squeeze main contrib non-free",
	"deb http://us.archive.ubuntu.com/ubuntu/ xenial main",
}

func TestLoadRepository_FileExists(t *testing.T) {
	repoListFile := newTestRepoListFile(t)
	defer os.Remove(repoListFile.Name())

	expected := repositoryTest
	actual := repoListAsStringSlice(LoadRepository(repoListFile.Name()))
	if !reflect.DeepEqual(expected, actual) {
		t.Fatal("expected != actual")
	}
}

func TestLoadRepository_DoesNotExist(t *testing.T) {
	repoListFile := newTestRepoListFile(t)
	if err := os.Remove(repoListFile.Name()); err != nil {
		t.Fatal(err)
	}

	actual := LoadRepository(repoListFile.Name())
	if actual == nil {
		t.Fatal("expected repoList not nil")
	}
	if len(actual) != 0 {
		t.Fatal("expected repoList length 0")
	}
}

func newTestRepoListFile(t *testing.T) *os.File {
	var contents string
	for _, s := range repositoryTest {
		contents += fmt.Sprintf("%s\n", s)
	}
	repoListFile, err := ioutil.TempFile("", "")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := repoListFile.Write([]byte(contents)); err != nil {
		t.Fatal(err)
	}
	if err := repoListFile.Close(); err != nil {
		t.Fatal(err)
	}
	return repoListFile
}

func repoListAsStringSlice(repoList debrepo.RepositoryList) []string {
	var repoStrings []string
	for _, r := range repoList {
		repoStrings = append(repoStrings, r.String())
	}
	return repoStrings
}
