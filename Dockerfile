FROM golang:1.23 AS builder

WORKDIR /go/src/app

# Copy go.mod and go.sum files first for dependency resolution
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy the rest of the application files
COPY cmd/services .

ENV GIN_MODE=release

# Build the Go application
RUN CGO_ENABLED=0 GOOS=linux go build -o main main.go

# Use a minimal base image for the final stage
FROM debian:bookworm-slim

# Install necessary packages
RUN apt-get update && DEBIAN_FRONTEND=noninteractive apt-get install -y \
    ca-certificates && \
    apt-get clean && \
    rm -rf /var/lib/apt/lists/*

WORKDIR /go/src/app

# Copy the binary from the builder stage
COPY --from=builder /go/src/app/main .

ENV GIN_MODE=release

# Expose the port the app runs on
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
