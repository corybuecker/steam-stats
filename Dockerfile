from golang:alpine
maintainer Cory Buecker <email@corybuecker.com>

run apk add --update --no-cache git
add . $GOPATH/src/github.com/corybuecker/steamfetcher
run go get github.com/corybuecker/steamfetcher
run go install github.com/corybuecker/steamfetcher
entrypoint ["/go/bin/steam-stats"]
