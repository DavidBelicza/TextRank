FROM golang:1.8
MAINTAINER David Belicza

ADD ./ /go/src/github.com/DavidBelicza/TextRank

WORKDIR /go/src/github.com/DavidBelicza/TextRank

CMD go get github.com/golang/dep/cmd/dep
CMD dep ensure
CMD /bin/bash