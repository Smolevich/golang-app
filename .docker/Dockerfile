FROM golang:1.13

ENV \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

RUN \
    go get -u -v github.com/golangci/golangci-lint/cmd/golangci-lint && \
    golangci-lint version

RUN go get -u -v .
WORKDIR /app

ADD . .

RUN go build -o app
