# running the file is required in the path../ to tighten dependencies in ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./auth_service/ ./auth_service
COPY ./internal/ ./internal
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR auth_service
RUN go build -o main .

EXPOSE 8102

CMD ["./main"]
