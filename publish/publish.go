package publish

import (
  "github.com/queeno/aptlify/mirror"
)

type PublishStruct struct {
  Name        string              `json:"name"`
  Resources   []ResourceStruct    `json:"resources"`
}

type ResourceStruct struct {
  Name        string                      `json:"name"`
  Type        string                      `json:"type"`
  Filter      []mirror.AptlyFilterStruct  `json:"filter"`
}
