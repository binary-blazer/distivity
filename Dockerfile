FROM golang:1.23.3-alpine

WORKDIR /app

COPY . .

RUN go build -o distivity main.go

ENTRYPOINT ["./distivity"]
