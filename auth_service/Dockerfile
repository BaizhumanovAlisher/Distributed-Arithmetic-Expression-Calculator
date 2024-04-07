# running the file is required in the path ../ to tighten dependencies in ../internal and ../src
# First stage: build the application
FROM golang:latest AS builder

WORKDIR /app

COPY ./auth_service/ ./auth_service
COPY ./internal/ ./internal
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR auth_service
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o main .

# Second stage: create the final image
FROM scratch

WORKDIR /app

COPY --from=builder /app/auth_service/main main
COPY ./auth_service/config.yaml config.yaml

EXPOSE 8102

CMD ["./main"]
