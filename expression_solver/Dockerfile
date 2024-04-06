# running the file is required in the path../ to tighten dependencies in ../proto, ../internal and ../go.work
FROM golang:latest

WORKDIR /app

COPY ./go.work ./go.work
COPY ./api_server/ ./api_server
COPY ./auth_service/ ./auth_service
COPY ./internal/ ./internal
COPY ./expression_solver/ ./expression_solver

WORKDIR expression_solver
RUN go build -o main .

EXPOSE 8103

CMD ["./main"]
