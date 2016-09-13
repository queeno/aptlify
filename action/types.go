package action

// Change types
const (
	Noop = iota
	MirrorUpdate
	MirrorCreate
	MirrorRecreate
	SnapshotUpdate
	RepoCreate
	GpgAdd
)

const (
	mirrorType = iota
	repoType
	gpgType
	snapshotType
)
