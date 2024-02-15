# Stage  1: Build the Go application
FROM golang:1.21.6

# Create a working directory
WORKDIR /app

# Copy go.mod and go.sum files
COPY go.mod go.sum ./

# Download all dependencies
RUN go mod download

# Copy the source code
COPY . .

# Copy the config.yaml file
COPY config.yaml /app/config.yaml

EXPOSE 8099

# Build the application
RUN CGO_ENABLED=0 go build -o /app/myapp

# Start the application
CMD ["/app/myapp"]
