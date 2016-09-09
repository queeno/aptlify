package config

import (
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/repo"
	"github.com/queeno/aptlify/gpg"
	"github.com/queeno/aptlify/snapshot"
	"encoding/json"
	"os"
)

type ConfigStruct struct {
	Mirrors 	[]mirror.AptlyMirrorStruct			`json:"mirrors"`
	Repos   	[]repo.AptlyRepoStruct					`json:"repos"`
	Gpg_keys	gpg.AptlyGpgStruct							`json:"gpg_keys"`
	Snapshots []snapshot.AptlySnapshotStruct	`json:"snapshots"`
}

var Config ConfigStruct = ConfigStruct{}
var State ConfigStruct = ConfigStruct{}

func (c *ConfigStruct) AddSnapshot(s snapshot.AptlySnapshotStruct){

	if s.IsEmpty() {
		return
	}

	c.Snapshots = append(c.Snapshots, s)
}

func (c *ConfigStruct) AddMirror(m mirror.AptlyMirrorStruct) {

	if m.IsEmpty() {
		return
	}

	c.Mirrors = append(c.Mirrors, m)
}

func (c *ConfigStruct) AddRepo(r repo.AptlyRepoStruct) {

	if r.IsEmpty() {
		return
	}

	c.Repos = append(c.Repos, r)
}


func (c *ConfigStruct) AddGpg(g string) {

	if g == "" {
		return
	}

	c.Gpg_keys.Fingerprint = append(c.Gpg_keys.Fingerprint, g)
}

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

func WriteConfig(filename string, config ConfigStruct) error {

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	encoded, err := json.MarshalIndent(config, "", "  ")
	if err != nil {
		return err
	}

	_, err = f.Write(encoded)

	return err

}
