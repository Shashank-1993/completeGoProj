# Step 1: Build the Go application
FROM golang:1.23.1 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source code into the container
COPY . .

# Build the Go app
RUN go build -o main .

# Step 2: Create a smaller image to run the application
FROM alpine:latest

# Set working directory in the second stage
WORKDIR /root/

# Copy the pre-built binary from the builder stage
COPY --from=builder /app/main .

# Expose port (optional, if your app listens on a port)
EXPOSE 8080

# Command to run the executable
CMD ["./main"]
