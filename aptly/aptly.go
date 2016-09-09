package aptly

import (
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/snapshot"
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
	// snapshot_create: IN: res ResourceStruct OUT: []string, error, string
	SnapshotCreate(snapshot.ResourceStruct) ([]string, error, string)
	// snapshot_filter: IN: ResourceStruct, string OUT: []string, error, string
	SnapshotFilter(snapshot.ResourceStruct, string) ([]string, error, string)
	// snapshot_drop: IN: snapname(string), force(bool), OUT: []string, err
	SnapshotDrop(string, bool) ([]string, error)
	// snapshot merge: IN []string snapshot names: OUT: []string, error, string (combined_snapshot_name)
	SnapshotMerge(string, []string) ([]string, error)
}
