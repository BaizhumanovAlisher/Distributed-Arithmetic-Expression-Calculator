# running the file is required in the path../ to tighten dependencies in ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./api_server/ ./api_server
COPY ./internal/ ./internal
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR api_server
RUN go build -o main .

EXPOSE 8099

CMD ["./main"]
