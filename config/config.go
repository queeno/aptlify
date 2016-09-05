package config

import (
	"github.com/queeno/aptlify/mirror"
	"github.com/queeno/aptlify/repo"
	"github.com/queeno/aptlify/gpg"
	"encoding/json"
	"os"
)

type ConfigStruct struct {
	Mirrors 	[]mirror.AptlyMirrorStruct	`json:"mirrors"`
	Repos   	[]repo.AptlyRepoStruct			`json:"repos"`
	Gpg_keys	gpg.AptlyGpgStruct					`json:"gpg_keys"`
}

var Config ConfigStruct = ConfigStruct{}
var State ConfigStruct = ConfigStruct{}

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

	f, err := os.Open(filename)

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
