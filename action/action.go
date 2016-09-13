package action

/* This should represent an atomic change to aptly

- State of change
- Resource to be changed
- Response
*/

import (
	"github.com/queeno/aptlify/config"
)

type ActionStruct struct {
	ResourceName     string
	ResourceType     int
	ChangeType       int
	changeReason     []string
	SnapshotRevision int
}

func (a ActionStruct) isEmpty() bool {
	if a.ResourceName == "" {
		return true
	}
	return false
}

func (a *ActionStruct) AddReasonToAction(reason string) {
	if reason == "" {
		a.changeReason = nil
	}
	a.changeReason = append(a.changeReason, reason)
}

func CreateActions(config *config.ConfigStruct, state *config.ConfigStruct) []ActionStruct {

	mirrorActions := createMirrorActions(config.Mirrors, state.Mirrors)
	repoActions := createRepoActions(config.Repos, state.Repos)
	gpgActions := createGpgActions(config.GpgKeys, state.GpgKeys)
	snapshotActions := createSnapshotActions(config.Snapshots, state.Snapshots)

	return append(append(append(gpgActions, mirrorActions...), repoActions...), snapshotActions...)
}
