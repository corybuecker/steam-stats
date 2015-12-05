package steam

import (
	"encoding/json"
	"fmt"

	"github.com/corybuecker/steam-stats/fetcher"
)

type Response struct {
	Games []OwnedGame `json:"games"`
}
type OwnedGame struct {
	ID              int    `json:"appid"`
	Name            string `json:"name"`
	PlaytimeForever int    `json:"playtime_forever"`
	PlaytimeRecent  int    `json:"playtime_2weeks"`
}

type OwnedGames struct {
	Response Response `json:"response"`
}

type Fetcher struct {
	SteamAPIKey string
	SteamID     string
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", fetcher.SteamAPIKey, fetcher.SteamID)
}

func (fetcher *Fetcher) GetOwnedGames(jsonfetcher fetcher.JSONFetcher) (*OwnedGames, error) {
	response, err := jsonfetcher.Fetch(fetcher.generateURL())
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
