# running the file is required in the path../ to tighten dependencies in ../proto, ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./go.work ./go.work
COPY ./auth_service/ ./auth_service
COPY ./protos/ ./protos
COPY ./internal/ ./internal
COPY ./api_server/ ./api_server

WORKDIR auth_service
RUN go mod download
RUN go build -o main .

EXPOSE 8102

CMD ["./main"]
