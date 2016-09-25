package actions

import (
	"github.com/corybuecker/wikipediasearch"
	"github.com/davecgh/go-spew/spew"
)

func SearchWikipedia(databaseHost string) error {
	results, _ := wikipediasearch.Search("Deus Ex", true)
	spew.Dump(results)
	return nil
}
