package steam

import (
	"fmt"
	"log"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/steam-stats-fetcher/database"
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
	SteamAPIKey string `bson:"steam_api_key"`
	SteamID     string `bson:"steam_id"`
	OwnedGames  OwnedGames
	Jsonfetcher jsonfetcher.Fetcher
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", fetcher.SteamAPIKey, fetcher.SteamID)
}

func (fetcher *Fetcher) GetOwnedGames() error {
	if err := fetcher.Jsonfetcher.Fetch(fetcher.generateURL(), &fetcher.OwnedGames); err != nil {
		return err
	}

	log.Printf("found %d games in the user's library", len(fetcher.OwnedGames.Response.Games))

	return nil
}
func (fetcher *Fetcher) UpdateOwnedGames(database database.Interface) error {
	for _, ownedGame := range fetcher.OwnedGames.Response.Games {
		ownedGameMap := map[string]interface{}{
			"id":              ownedGame.ID,
			"name":            ownedGame.Name,
			"playtimeForever": ownedGame.PlaytimeForever,
			"playtimeRecent":  ownedGame.PlaytimeRecent,
		}

		log.Printf("upserting %s games in the user's library", ownedGame.Name)
		if err := database.Upsert(map[string]interface{}{"id": ownedGame.ID}, ownedGameMap); err != nil {
			return err
		}
	}
	return nil
}
