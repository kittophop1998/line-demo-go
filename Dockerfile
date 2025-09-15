# ---------------------------------------------------------
# Stage 1: Build Go binary
# ---------------------------------------------------------
FROM golang:1.23.8-alpine AS builder

# Set workdir
WORKDIR /app

# Copy go.mod/go.sum and download dependencies
COPY go.mod go.sum ./
RUN go mod download && go mod verify

# Copy project files
COPY . .

# Move to working dir for main.go
WORKDIR /app/cmd/server

# Build Go binary
RUN go build -o server .

# ---------------------------------------------------------
# Stage 2: Run with Google Chrome
# ---------------------------------------------------------
FROM debian:stable-slim

# Set workdir
WORKDIR /root/

# Set environment variable
ENV ENV_NAME=sit

# Copy binary from builder
COPY --from=builder /app/cmd/server/server .

# (Optional) Copy configs if needed
COPY configs/ /root/configs/

# Expose port
EXPOSE 8080

# Command
CMD ["./server"]