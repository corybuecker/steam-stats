package fetcher

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type Interface interface {
	Fetch(string, interface{}) error
}

type JSONFetcher struct{}

func (fetcher *JSONFetcher) Fetch(url string, structToLoad interface{}) error {
	var err error
	response, err := http.Get(url)

	if err != nil {
		return err
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return err
	}

	err = json.Unmarshal(contents, structToLoad)

	if err != nil {
		return err
	}

	return nil
}
