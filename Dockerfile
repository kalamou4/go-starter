FROM golang:1.21-alpine

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN go build -o main cmd/api/main.go

# Expose port
EXPOSE 8080

# Run the application
CMD ["./main"]
