package config

import (
  "os"
  "encoding/json"
)

type AptlyFilterStruct struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type AptlyMirrorStruct struct {
	Name       string              `json:"name"`
	Gpg        []string            `json:"gpg"`
	Url        string              `json:"url"`
	Dist       string              `json:"dist"`
	Component  string              `json:"component"`
	Filter     []AptlyFilterStruct `json:"filter"`
	FilterDeps bool                `json:"filter-with-deps"`
}

type ConfigStruct struct {
	Mirrors []AptlyMirrorStruct		`json:"mirrors"`
	Repos   []string							`json:"repos"`
}

var Config ConfigStruct

// Open configuration file and decode the JSON
func LoadConfig(filename string, config *ConfigStruct) error {

  f, err := os.Open(filename)

  if err != nil {
    return err
  }

  defer f.Close()

  dec := json.NewDecoder(f)

  err = dec.Decode(config)
  if err != nil {
    return err
  }

  return nil

}
