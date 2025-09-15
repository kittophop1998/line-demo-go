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
# Stage 2: Run
# ---------------------------------------------------------
FROM alpine:latest

# Install ca-certificates for TLS verification
RUN apk --no-cache add ca-certificates tzdata

ENV ENV_NAME=dev

# Set workdir
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/cmd/server/server .

# (Optional) Copy configs if needed
COPY configs ./configs

# Expose port
EXPOSE 8080

# Command to run
CMD ["./server"]
