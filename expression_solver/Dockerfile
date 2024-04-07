# running the file is required in the path ../ to tighten dependencies in ../internal and ../src
# First stage: build the application
FROM golang:latest AS builder

WORKDIR /app

COPY ./expression_solver/ ./expression_solver
COPY ./internal/ ./internal
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR expression_solver
RUN CGO_ENABLED=0 go build -ldflags '-extldflags "-static"' -o main .

# Second stage: create the final image
FROM scratch

WORKDIR /app

COPY --from=builder /app/expression_solver/main main
COPY ./expression_solver/config.yaml config.yaml

EXPOSE 8103

CMD ["./main"]
