FROM golang:1.13.7-buster

WORKDIR /go/src/app

ENV GO111MODULE=on

RUN apt-get update -qq && \
    apt-get install -y git

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN go get bitbucket.org/liamstask/goose/cmd/goose

CMD ["/bin/bash"]
