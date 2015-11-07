package configuration

import (
	"encoding/json"
	"os"
)

type Configuration struct {
	SteamAPIKey     string
	SteamID         string
	GiantBombAPIKey string
}

func (configuration *Configuration) Load(filePath string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	decoder := json.NewDecoder(file)
	if err = decoder.Decode(configuration); err != nil {
		return err
	}
	return nil
}
