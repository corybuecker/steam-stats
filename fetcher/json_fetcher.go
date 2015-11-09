package fetcher

import (
	"io/ioutil"
	"net/http"
)

type JSONFetcherInterface interface {
	fetch(url string) ([]byte, error)
}

type JSONFetcher struct{}

func (fetcher *JSONFetcher) fetch(url string) ([]byte, error) {
	response, err := http.Get(url)

	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)

	if err != nil {
		return nil, err
	}

	return contents, nil
}
