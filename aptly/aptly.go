package aptly

import (
	"github.com/queeno/aptlify/mirror"
)

type Aptly interface {
	// mirror_create: IN: mirror struct OUT: error
	Mirror_create(mirror.AptlyMirrorStruct) ([]string, error)
	// mirror_list: IN: n/a OUT: mirror-list, error
	Mirror_list() ([]string, error)
	// mirror_update: IN: mirror-name OUT: error
	Mirror_update(string) ([]string, error)
	// repo_list: IN: n/a OUT: repo-list, error
	Repo_list() ([]string, error)
	// repo_add: IN: repo-name, OUT: error
	Repo_create(string) ([]string, error)
}
