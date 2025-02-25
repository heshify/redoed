# Use the official Golang image
FROM golang:1.24-alpine

# Set working directory inside the container
WORKDIR /app

# Copy everything into the container
COPY . .

# Build the Go application
RUN go build -o server ./cmd/server

# Expose port 8080
EXPOSE 8080

# Run the binary
CMD ["./server"]
