package configuration

import "github.com/corybuecker/steam-stats/database"

type Configuration struct {
	SteamAPIKey     string
	SteamID         string
	GiantBombAPIKey string
}

func (configuration *Configuration) Load(database database.Interface) error {
	var row map[string]interface{}
	var err error

	row, err = database.GetRow("configurations", "steam_stats", "configuration")

	if err != nil {
		return err
	}

	configuration.SteamAPIKey = row["steamApiKey"].(string)
	configuration.SteamID = row["steamId"].(string)
	configuration.GiantBombAPIKey = row["giantBombApiKey"].(string)

	return nil
}
