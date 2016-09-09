package snapshot

import (
  "github.com/queeno/aptlify/mirror"
)

type AptlySnapshotStruct struct {
  Name        string              `json:"name"`
  Resources   []ResourceStruct    `json:"resources"`
  Revision    int                 `json:"revision"`
}

type ResourceStruct struct {
  Name        string                      `json:"name"`
  Type        string                      `json:"type"`
  Filter      []mirror.AptlyFilterStruct  `json:"filter"`
}


func (s AptlySnapshotStruct) IsEmpty() bool {

  if s.Name == "" {
    return true
  }

  return false
}

func (thisSnapshot AptlySnapshotStruct)SearchSnapshotInAptlySnapshots(snapshots []AptlySnapshotStruct) AptlySnapshotStruct {

  for _, snapshot := range snapshots {
    if snapshot.Name == thisSnapshot.Name {
      return snapshot
    }
  }

  return AptlySnapshotStruct{}

}
