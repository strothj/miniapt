package miniapt

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/strothj/go-debrepo/debrepo"
)

// LoadRepository reads a list of repositories from the specified file. Any line
// not in the format
// "deb http://ftp.debian.org/debian squeeze main contrib non-free" is ignored.
func LoadRepository(repoListPath string) (repoList debrepo.RepositoryList) {
	b, err := ioutil.ReadFile(repoListPath)
	if err != nil {
		return
	}
	for _, l := range strings.Split(string(b), "\n") {
		if len(l) > 0 {
			source, err := debrepo.ParseRepository(l)
			if err == nil {
				repoList = append(repoList, source)
			}
		}
	}
	return
}

// SaveRepositories saves a repository source list.
func SaveRepositories(repoListPath string, repoList debrepo.RepositoryList) error {
	var b string
	for _, l := range repoList {
		b += fmt.Sprintf("%s\n", l)
	}
	return ioutil.WriteFile(repoListPath, []byte(b), 0755)
}
