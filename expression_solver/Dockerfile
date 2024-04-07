# running the file is required in the path../ to tighten dependencies in ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./internal/ ./internal
COPY ./expression_solver/ ./expression_solver
COPY ./src/ ./src

ENV GOPATH=/app/src
WORKDIR expression_solver
RUN go build -o main .

EXPOSE 8103

CMD ["./main"]
