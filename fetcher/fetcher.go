package fetcher

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/corybuecker/steam-stats/database"
)

type mySteamGames struct {
	Response struct {
		Games []struct {
			Appid           int `json:"appid"`
			PlaytimeForever int `json:"playtime_forever"`
		} `json:"games"`
	} `json:"response"`
}

type Fetcher struct {
	Storage         *database.DB
	SteamAPIKey     string
	SteamID         string
	GiantBombAPIKey string
	client          *http.Client
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json", fetcher.SteamAPIKey, fetcher.SteamID)
}

func (fetcher *Fetcher) FetchAll() mySteamGames {
	if fetcher.client == nil {
		fetcher.client = &http.Client{}
	}
	return fetch(fetcher.generateURL())
}

func fetch(url string) mySteamGames {
	response, _ := http.Get(url)

	defer response.Body.Close()
	contents, _ := ioutil.ReadAll(response.Body)
	var games mySteamGames
	json.Unmarshal(contents, &games)

	return games
}
