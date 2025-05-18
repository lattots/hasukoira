# Use Go version 1.24.3 on Alpine OS
FROM golang:1.24.3-alpine

# Set the working directory inside the container
WORKDIR /app

# Download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source files
COPY ./cmd ./cmd
COPY ./pkg ./pkg

# Build runnable binary from source
RUN go build -o main ./cmd/main.go

COPY ./data ./data

# Run the built binary (looks for ./main in /app)
CMD ["./main"]
