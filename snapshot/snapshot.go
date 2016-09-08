package snapshot

import (
  "github.com/queeno/aptlify/mirror"
)

type SnapshotStruct struct {
  Name        string              `json:"name"`
  Resources   []ResourceStruct    `json:"resources"`
  Revision    int                 `json:"revision"`
}

type ResourceStruct struct {
  Name        string                      `json:"name"`
  Type        string                      `json:"type"`
  Filter      []mirror.AptlyFilterStruct  `json:"filter"`
}


func (s SnapshotStruct) IsEmpty() bool {

  if s.Name == "" {
    return true
  }

  return false
}

func (thisSnapshot SnapshotStruct)SearchSnapshotInAptlySnapshots(snapshots []SnapshotStruct) SnapshotStruct {

  for _, snapshot := range snapshots {
    if snapshot.Name == thisSnapshot.Name {
      return snapshot
    }
  }

  return SnapshotStruct{}

}
