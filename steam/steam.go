package steam

import (
	"errors"
	"fmt"
	"log"

	"github.com/corybuecker/jsonfetcher"
	"github.com/corybuecker/mgoconfig"
	"github.com/corybuecker/steamfetcher/database"
)

type ownedGames struct {
	Response struct {
		Games []database.Game `json:"games"`
	} `json:"response"`
}

type configuration struct {
	SteamAPIKey string `bson:"steam_api_key"`
	SteamID     string `bson:"steam_id"`
}

type Fetcher struct {
	ConfigurationSettings *mgoconfig.Configuration
	OwnedGames            ownedGames
	jsonFetcher           jsonfetcher.Fetcher
	configuration         configuration
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("https://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", fetcher.configuration.SteamAPIKey, fetcher.configuration.SteamID)
}

func (fetcher *Fetcher) configure(database database.Interface) error {
	if fetcher.ConfigurationSettings == nil {
		fetcher.ConfigurationSettings = &mgoconfig.Configuration{
			Database: "steamfetcher",
			Key:      "steam",
			Session:  database.GetSession(),
		}
	}

	if err := fetcher.ConfigurationSettings.Get(&fetcher.configuration); err != nil {
		if err.Error() == "not found" {
			return errors.New("the steam configuration could not be fetched")
		}
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
