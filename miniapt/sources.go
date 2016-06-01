package miniapt

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/strothj/debrepo"
)

// LoadSources parses a repository source list.
func LoadSources(sourcelistPath string) (sourceList debrepo.SourceList) {
	b, err := ioutil.ReadFile(sourcelistPath)
	if err != nil {
		return
	}
	for _, l := range strings.Split(string(b), "\n") {
		if len(l) > 0 {
			source, err := debrepo.ParseSource(l)
			if err == nil {
				sourceList = append(sourceList, source)
			}
		}
	}
	return
}

// SaveSources saves a repository source list.
func SaveSources(sourcelistPath string, sourceList debrepo.SourceList) error {
	var b string
	for _, l := range sourceList {
		b += fmt.Sprintf("%s\n", l)
	}
	return ioutil.WriteFile(sourcelistPath, []byte(b), 0755)
}
