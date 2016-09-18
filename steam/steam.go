package steam

import (
	"fmt"
	"log"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/mgoconfig"
	"github.com/corybuecker/steam-stats-fetcher/database"
)

type ownedGame struct {
	ID              int    `json:"appid" bson:"steam_id"`
	Name            string `json:"name" bson:"name"`
	PlaytimeForever int    `json:"playtime_forever" bson:"playtimeForever"`
	PlaytimeRecent  int    `json:"playtime_2weeks" bson:"playtimeRecent"`
}

type ownedGames struct {
	Response struct {
		Games []ownedGame `json:"games"`
	} `json:"response"`
}
type configuration struct {
	SteamAPIKey string `bson:"steam_api_key"`
	SteamID     string `bson:"steam_id"`
}
type Fetcher struct {
	configuration configuration
	OwnedGames    ownedGames
	jsonFetcher   jsonfetcher.Fetcher
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", fetcher.configuration.SteamAPIKey, fetcher.configuration.SteamID)
}

func (fetcher *Fetcher) configure(database database.Interface) error {
	if err := mgoconfig.Get(database.GetSession(), "steam_stats_fetcher", "steam", &fetcher.configuration); err != nil {
		return err
	}
	fetcher.jsonFetcher = &jsonfetcher.Jsonfetcher{}

	return nil
}

func (fetcher *Fetcher) getOwnedGames() error {
	if err := fetcher.jsonFetcher.Fetch(fetcher.generateURL(), &fetcher.OwnedGames); err != nil {
		return err
	}

	log.Printf("found %d games in the user's library", len(fetcher.OwnedGames.Response.Games))

	return nil
}
func (fetcher *Fetcher) UpdateOwnedGames(database database.Interface) error {
	var err error

	if err = fetcher.configure(database); err != nil {
		return err
	}

	if err = fetcher.getOwnedGames(); err != nil {
		return err
	}

	for _, ownedGame := range fetcher.OwnedGames.Response.Games {
		log.Printf("upserting %s games in the user's library", ownedGame.Name)
		if err = database.UpsertIntField("steam_id", ownedGame.ID, ownedGame); err != nil {
			return err
		}
	}
	return nil
}
