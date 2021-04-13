FROM golang:1.15
MAINTAINER David Belicza

ADD ./ /go/src/github.com/DavidBelicza/TextRank

WORKDIR /go/src/github.com/DavidBelicza/TextRank

CMD go mod vendor
CMD /bin/bash
