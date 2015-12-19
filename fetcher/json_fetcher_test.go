package fetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server
var sampleResponse string = "{\"response\": {\"games\": [{\"appid\": 10, \"playtime_forever\": 32}]}}"

type MarshalStruct struct {
	Response struct {
		Games []struct {
			ID              int    `json:"appid"`
			Name            string `json:"name"`
			PlaytimeForever int    `json:"playtime_forever"`
			PlaytimeRecent  int    `json:"playtime_2weeks"`
		} `json:"games"`
	} `json:"response"`
}

type BadMarshalStruct struct {
	Response struct {
	} `json:"bad"`
}

func buildServer() {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, sampleResponse)
	}))
}

func init() {
	buildServer()
}

func TestSuccessfulMarshal(t *testing.T) {
	var fetcher JSONFetcher = JSONFetcher{}
	data := MarshalStruct{}
	fetcher.Fetch(server.URL, &data)
	if data.Response.Games[0].ID != 10 {
		t.Error("expected to return sample response")
	}
}

func TestUnsuccessfulMarshal(t *testing.T) {
	var fetcher JSONFetcher = JSONFetcher{}
	data := BadMarshalStruct{}
	emptyStruct := BadMarshalStruct{}
	fetcher.Fetch(server.URL, &data)
	if data != emptyStruct {
		t.Error("expected an empty struct")
	}
}
