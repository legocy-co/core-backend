FROM golang:1.20-alpine AS builder

RUN go version

COPY . /src/app
WORKDIR /src/app/

RUN go mod download
RUN GOOS=linux go build -o ./.bin/app ./cmd/legocy/main.go

FROM alpine:latest

WORKDIR /root/

COPY --from=0 /src/app/.bin/app .
COPY --from=0 /src/app/config ./config

CMD ["./app"]