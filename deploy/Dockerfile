FROM golang:1.20

WORKDIR /src/app

COPY ../ ./

RUN go mod download

WORKDIR /src/app/cmd/legocy

RUN go build -o /legocy_bin

EXPOSE 8080

CMD ["/legocy_bin"]
