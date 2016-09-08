package action

import (
  "github.com/queeno/aptlify/snapshot"
)


func createSnapshotActions(configSnapshots []snapshot.SnapshotStruct, stateSnapshots []snapshot.SnapshotStruct) []ActionStruct {

	var actions = []ActionStruct{}

	for _, configSnapshot := range configSnapshots {

		stateSnapshot := configSnapshot.SearchSnapshotInAptlySnapshots(stateSnapshots)
		if stateSnapshot.IsEmpty() {
			actions = append(actions, updateSnapshot(configSnapshot, 0))
		} else {
			actions = append(actions, updateSnapshot(configSnapshot, stateSnapshot.Revision))
		}
	}

	return actions

}

func updateSnapshot (a snapshot.SnapshotStruct, revision int) ActionStruct {

	var ac = ActionStruct{	ResourceName: a.Name,
													ChangeType: Snapshot_update,
													ResourceType: snapshotType,
												 	SnapshotRevision: revision+1 }

	ac.AddReasonToAction("update_snapshot")

	return ac

}
