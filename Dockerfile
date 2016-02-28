from golang:latest
maintainer Cory Buecker <email@corybuecker.com>

add . $GOPATH/src/github.com/corybuecker/steam-stats-fetcher
run go get github.com/corybuecker/steam-stats-fetcher
run go install github.com/corybuecker/steam-stats-fetcher
entrypoint ["/go/bin/steam-stats-fetcher"]
