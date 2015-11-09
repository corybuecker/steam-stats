package fetcher

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

var server *httptest.Server
var sampleResponse string = "{\"response\": {\"games\": [{\"appid\": 10, \"playtime_forever\": 32}]}}"

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

func TestSuccessfulCall(t *testing.T) {
	var fetcher *JSONFetcher = new(JSONFetcher)
	response, _ := fetcher.fetch(server.URL)
	if bytes.Compare(response, []byte(sampleResponse)) != 0 {
		t.Error("expected to return sample response")
	}
}
