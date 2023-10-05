FROM golang:1.20

WORKDIR /src/app

COPY ../ ./

RUN go mod download

WORKDIR /src/app/cmd/legocy

RUN CGO_ENABLED=0 GOOS=linux go build -o /legocy_bin

WORKDIR /src/app

EXPOSE 8080

CMD ["/cmd/legocy/legocy_bin"]
