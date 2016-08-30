package utils

import (
  "os"
  "encoding/json"
)

type ConfigStruct struct {
  captain_host  string    `json:captain_host`
  docktor_host  string    `json:docktor_host`
}

var Config = ConfigStruct {
  captain_host: "",
  docktor_host: "",
}

// Open configuration file and decode the JSON
func LoadConfig(filename string, config *ConfigStruct) error {

  f, err := os.Open(filename)

  if err != nil {
    return err
  }

  defer f.Close()

  dec := json.NewDecoder(f)

  err = dec.Decode(Config)
  if err != nil {
    return err
  }

  return nil

}
