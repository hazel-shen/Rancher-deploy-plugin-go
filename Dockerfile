FROM golang:alpine
WORKDIR /
ADD . /
RUN apk add --no-cache git mercurial \
    && go get github.com/tidwall/sjson \
    && apk del git mercurial
RUN cd / && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main
ENTRYPOINT ["/main"]