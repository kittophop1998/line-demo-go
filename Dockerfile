# ===== Stage 1: Build =====
FROM golang:1.21-alpine AS builder

# Set working directory
WORKDIR /app

# Copy go.mod and go.sum first (to cache deps)
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build binary
RUN go build -o bot .

# ===== Stage 2: Run =====
FROM alpine:latest

# Install ca-certificates (required for HTTPS requests)
RUN apk --no-cache add ca-certificates

# Set working directory
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/bot .

# Set environment variables (optional)
# ENV LINE_CHANNEL_SECRET=your_secret
# ENV LINE_CHANNEL_TOKEN=your_token
# ENV PORT=8080

# Expose port
EXPOSE 8080

# Command to run
CMD ["./bot"]
