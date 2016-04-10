from golang:alpine
maintainer Cory Buecker <email@corybuecker.com>

run apk add --update --no-cache git
add . $GOPATH/src/github.com/corybuecker/steam-stats-fetcher
run go get github.com/corybuecker/steam-stats-fetcher
run go install github.com/corybuecker/steam-stats-fetcher
entrypoint ["/go/bin/steam-stats"]
