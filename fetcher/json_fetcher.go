package fetcher

import (
	"encoding/json"
	"errors"
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

	if response.StatusCode != 200 {
		return errors.New("the HTTP call returned a non-200 response")
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
