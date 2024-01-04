FROM golang:1.20 AS builder

WORKDIR /src/app/

COPY go.* ./
RUN go mod download

COPY . /src/app
RUN go build -o ./.bin/app ./cmd/legocy/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /src/app/.bin/app .
COPY --from=0 /src/app/config ./config

CMD ["./app"]