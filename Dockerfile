# Stage 1: Build Go app
FROM golang:1.21-alpine AS builder
WORKDIR /app

# Copy go.mod and go.sum to download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the entire source code
COPY . .

# Build the Go app with specific environment variables
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o main ./main.go

# Stage 2: Create a smaller image for running the app
FROM alpine:latest
WORKDIR /app

# Copy the compiled binary and .env file
COPY --from=builder /app/main .
COPY .env .env

# Expose the application port
EXPOSE 8080

# Run the app
CMD ["./main"]

