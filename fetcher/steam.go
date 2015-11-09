package fetcher

import (
	"encoding/json"
	"fmt"

	"github.com/corybuecker/steam-stats/database"
)

type OwnedGames struct {
	Response struct {
		Games []struct {
			Appid           int `json:"appid"`
			PlaytimeForever int `json:"playtime_forever"`
		} `json:"games"`
	} `json:"response"`
}

type SteamFetcher struct {
	Storage         *database.DB
	SteamAPIKey     string
	SteamID         string
	GiantBombAPIKey string
}

func (fetcher *SteamFetcher) generateURL() string {
	return fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", fetcher.SteamAPIKey, fetcher.SteamID)
}

func (fetcher *SteamFetcher) GetOwnedGames(jsonfetcher JSONFetcher) (*OwnedGames, error) {
	response, err := jsonfetcher.fetch(fetcher.generateURL())
	if err != nil {
		return nil, err
	}

	var ownedGames = new(OwnedGames)

	err = json.Unmarshal(response, ownedGames)

	if err != nil {
		return nil, err
	}

	return ownedGames, nil
}
