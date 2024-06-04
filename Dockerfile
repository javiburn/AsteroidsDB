# Stage 1: Build the Go binary
FROM golang:1.22.3 AS builder

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go.mod and go.sum files first to leverage Docker cache
COPY go.mod go.sum ./

# Copy the .env file if it exists
COPY .env ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the remaining source code into the container
COPY . .

# Build the Go app
RUN CGO_ENABLED=0 GOOS=linux go build -v -a -installsuffix cgo -o app ./cmd

# Stage 2: Create a minimal image
FROM alpine:latest

# Set the working directory
WORKDIR /app

# Copy the binary and .env from the builder stage
COPY --from=builder /app/app /app/
COPY --from=builder /app/.env /app/

# Expose port 8080
EXPOSE 8080

# Define environment variable
ENV PORT=8080

# Command to run the executable
ENTRYPOINT ["/app/app"]
