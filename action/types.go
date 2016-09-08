package action

// Change types
const (
	Noop = iota
	Mirror_update
	Mirror_create
	Mirror_recreate
	Snapshot_update
	Repo_create
	Gpg_add
)

const (
	mirrorType = iota
	repoType
	gpgType
	snapshotType
)
