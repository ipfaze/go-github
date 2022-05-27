FROM golang:1.17
MAINTAINER Léo Unbekandt "leo@scalingo.com"

RUN go get github.com/cespare/reflex

WORKDIR $GOPATH/src/github.com/Scalingo/sclng-backend-test-v1

EXPOSE 5000

CMD $GOPATH/bin/sclng-backend-test-v1

