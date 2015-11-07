package fetcher

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server

func buildServer(headers http.Header) {
	server = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		for header, value := range headers {
			w.Header().Set(header, value[0])
		}
		w.Header().Add("Content-Type", "application/json")
		w.WriteHeader(200)
		fmt.Fprint(w, "{\"response\": {\"games\": [{\"appid\": 10, \"playtime_forever\": 32}]}}")
	}))
}

func init() {
	headers := http.Header{}
	buildServer(headers)
}

func TestCreatedClient(t *testing.T) {
	fetcher := new(Fetcher)
	fetcher.FetchAll()
	if fetcher.client == nil {
		t.Fatalf("should have created a parser")
	}
}

func TestFetch(t *testing.T) {
	games := fetch(server.URL)
	if games.Response.Games[0].Appid != 10 {
		t.Error("expected data to be parsed")
	}
}
