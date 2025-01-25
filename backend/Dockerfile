# Backend Dockerfile
FROM golang:1.18 AS builder

# Set working directory
WORKDIR /app

# Copy Go modules and dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application code
COPY . .

# Build the Go app
RUN go build -o main .

# Use a minimal image for running the application
FROM alpine:3.17

# Install required dependencies
RUN apk --no-cache add ca-certificates

# Set working directory and copy built binary
WORKDIR /root/
COPY --from=builder /app/main .

# Expose the backend port
EXPOSE 8080

# Command to run the application
CMD ["./main"]
