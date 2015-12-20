from golang:latest
maintainer Cory Buecker <email@corybuecker.com>

add . $GOPATH/src/github.com/corybuecker/steam-stats
run go get github.com/corybuecker/steam-stats
run go install github.com/corybuecker/steam-stats
entrypoint /go/bin/steam-stats
