package jobs

import (
	"log"

	"github.com/corybuecker/steam-stats-fetcher/database"
	"github.com/corybuecker/steam-stats-fetcher/fetcher"
	"github.com/corybuecker/steam-stats-fetcher/giantbomb"
	"github.com/corybuecker/steam-stats-fetcher/steam"
)

type Job struct {
	Fetcher  *fetcher.JSONFetcher
	Database database.Interface
}

func (job *Job) OwnedGamesFetch(steamFetcher *steam.Fetcher) {
	if err := steamFetcher.GetOwnedGames(job.Fetcher); err != nil {
		log.Fatalln(err.Error())
	}

	if err := steamFetcher.UpdateOwnedGames(job.Database); err != nil {
		log.Fatalln(err.Error())
	}
}

func (job *Job) OwnedGamesSearch(steamFetcher *steam.Fetcher, giantBombFetcher *giantbomb.Fetcher) {
	var ownedGamesWithoutGiantBomb map[int]string
	var err error

	if ownedGamesWithoutGiantBomb, err = steamFetcher.FetchOwnedGamesWithoutGiantBomb(job.Database); err != nil {
		log.Fatalln(err.Error())
	}

	for ownedGameId, ownedGameName := range ownedGamesWithoutGiantBomb {
		if err := giantBombFetcher.FindOwnedGame(job.Fetcher, ownedGameName); err != nil {
			log.Println(err.Error())
		}
		if err := giantBombFetcher.UpdateFoundGames(ownedGameId, job.Database); err != nil {
			log.Println(err.Error())
		}
	}
}

func (job *Job) OwnedGamesFetchByID(steamFetcher *steam.Fetcher, giantBombFetcher *giantbomb.Fetcher) {
	var ownedGamesWithGiantBombID map[int]int
	var err error

	if ownedGamesWithGiantBombID, err = steamFetcher.FetchOwnedGamesGiantBombID(job.Database); err != nil {
		log.Fatalln(err.Error())
	}

	for ownedGameId, giantbombId := range ownedGamesWithGiantBombID {
		if err := giantBombFetcher.FindGameByID(job.Fetcher, giantbombId); err != nil {
			log.Println(err.Error())
		}
		if err := giantBombFetcher.UpdateFoundGames(ownedGameId, job.Database); err != nil {
			log.Println(err.Error())
		}
	}
}
