package steam

import (
	"fmt"
	"log"

	"github.com/corybuecker/steam-stats/database"
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
	OwnedGames  OwnedGames
}

func (fetcher *Fetcher) generateURL() string {
	return fmt.Sprintf("http://api.steampowered.com/IPlayerService/GetOwnedGames/v0001/?key=%s&steamid=%s&format=json&include_appinfo=1&include_played_free_games=1", fetcher.SteamAPIKey, fetcher.SteamID)
}

func (fetcher *Fetcher) GetOwnedGames(jsonfetcher fetcher.Interface) error {
	if err := jsonfetcher.Fetch(fetcher.generateURL(), &fetcher.OwnedGames); err != nil {
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

		if err := database.Upsert("videogames", "ownedgames", ownedGameMap); err != nil {
			return err
		}
	}
	return nil
}

func (fetcher *Fetcher) FetchOwnedGamesWithoutGiantBomb(database database.Interface) ([]string, error) {
	var gamesList []map[string]interface{}
	var err error

	var games = make([]string, 0)

	if gamesList, err = database.RowsWithoutFields("videogames", "ownedgames", []string{"giantbomb", "giantbomb_id"}); err != nil {
		return nil, err
	}

	for _, game := range gamesList {
		games = append(games, game["name"].(string))
	}

	return games, nil
}

func (fetcher *Fetcher) FetchOwnedGamesGiantBombID(database database.Interface) ([]int, error) {
	var gamesList []map[string]interface{}
	var err error

	var games = make([]int, 0)

	if gamesList, err = database.RowsWithField("videogames", "ownedgames", "giantbomb_id"); err != nil {
		return nil, err
	}

	for _, game := range gamesList {
		games = append(games, int(game["giantbomb_id"].(float64)))
	}

	return games, nil
}
