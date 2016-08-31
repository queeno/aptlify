package aptly

import (
	"github.com/queeno/aptlify/config"
)

type Aptly interface {
		// mirror_create: IN: name, url, dist, component, filter, filter-with-deps, OUT: error
		Mirror_create(string, string, string, string, []AptlyFilterStruct, bool) error
		// mirror_list: IN: n/a OUT: mirror-list, error
    Mirror_list() ([]string, error)
		// mirror_update: IN: mirror-name OUT: error
		Mirror_update(string) error
		// repo_list: IN: n/a OUT: repo-list, error
		Mepo_list() ([]string, error)
		// repo_add: IN: repo-name, OUT: error
		Repo_add(string) error
}
