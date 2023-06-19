# Use golang version 1.19
FROM golang:1.19-alpine

# Define working directory
WORKDIR /app

# Copy go.mod and go sum to workdir
COPY go.mod go.sum /app/

# Install dependencies
RUN go mod download

# Copy rest of file to workdir
COPY . .

# Build application
RUN go build -o main

# Expose port 8000
EXPOSE 8000

# Run main when container running
CMD ["./main"]