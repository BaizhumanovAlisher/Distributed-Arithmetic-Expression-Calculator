# running the file is required in the path../ to tighten dependencies in ../proto, ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./go.work ./go.work
COPY ./api_server/ ./api_server
COPY ./protos/ ./protos
COPY ./internal/ ./internal
COPY ./auth_service/ ./auth_service
COPY ./expression_solver/ ./expression_solver

WORKDIR api_server
RUN go build -o main .

EXPOSE 8099

CMD ["./main"]
