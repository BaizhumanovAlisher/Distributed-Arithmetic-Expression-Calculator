# running the file is required in the path ../ to tighten dependencies in ../internal and ../src
# First stage: build the application
FROM golang:latest AS builder

WORKDIR /app

COPY ./api_server/ ./api_server
COPY ./internal/ ./internal
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR api_server
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o main .

# Second stage: create the final image
FROM scratch

WORKDIR /app

COPY --from=builder /app/api_server/main main
COPY ./api_server/config.yaml config.yaml

EXPOSE 8099

CMD ["./main"]
