package jobs

import (
	"log"

	"github.com/corybuecker/steam-stats/database"
	"github.com/corybuecker/steam-stats/fetcher"
	"github.com/corybuecker/steam-stats/giantbomb"
	"github.com/corybuecker/steam-stats/steam"
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
	var ownedGamesWithoutGiantBomb []string
	var err error

	if ownedGamesWithoutGiantBomb, err = steamFetcher.FetchOwnedGamesWithoutGiantBomb(job.Database); err != nil {
		log.Fatalln(err.Error())
	}

	for _, ownedGame := range ownedGamesWithoutGiantBomb {
		if err := giantBombFetcher.FindOwnedGame(job.Fetcher, ownedGame); err != nil {
			log.Println(err.Error())
		}
		if err := giantBombFetcher.UpdateFoundGames(job.Database); err != nil {
			log.Println(err.Error())
		}
	}
}

func (job *Job) OwnedGamesFetchByID(steamFetcher *steam.Fetcher, giantBombFetcher *giantbomb.Fetcher) {
	var ownedGamesWithGiantBombID []int
	var err error

	if ownedGamesWithGiantBombID, err = steamFetcher.FetchOwnedGamesGiantBombID(job.Database); err != nil {
		log.Fatalln(err.Error())
	}

	for _, ownedGame := range ownedGamesWithGiantBombID {
		if err := giantBombFetcher.FindGameByID(job.Fetcher, ownedGame); err != nil {
			log.Println(err.Error())
		}
		if err := giantBombFetcher.UpdateFoundGames(job.Database); err != nil {
			log.Println(err.Error())
		}
	}
}
