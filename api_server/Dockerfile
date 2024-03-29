FROM golang:latest

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .
COPY config.yaml ./config.yaml

RUN go build -o main .

EXPOSE 8099

CMD ["./main"]
