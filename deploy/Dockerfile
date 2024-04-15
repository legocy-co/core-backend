FROM golang:1.20-bookworm AS builder

WORKDIR /src/app

COPY . .

COPY .env ./

RUN go mod download
RUN go mod verify
RUN go build -o ./.bin/app ./cmd/legocy/main.go


EXPOSE 8080
CMD ["./.bin/app"]
