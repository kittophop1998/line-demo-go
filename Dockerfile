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

ENV LINE_CHANNEL_SECRET=3905d4f46c24c2475a877125cd81c748 
ENV LINE_CHANNEL_ACCESS_TOKEN=qzzM//7rjcjHKHmBcWcHDe0BawtyuiOHekg4i0HHdu+UVdb5P0WZwIAGqK8ULDFwcA+1jbUYBf0EuduAtgjoD1Ejj4wcdveVZbXQAyhMl/1MfkvY9o64o4T9p9x03Tksk54bvujGn99jg1AbU42FNAdB04t89/1O/w1cDnyilFU=

# Set workdir
WORKDIR /app

# Copy binary from builder
COPY --from=builder /app/cmd/server/server .

# Expose port
EXPOSE 8080

# Use environment variables at runtime instead of hardcoding
# LINE_CHANNEL_SECRET, LINE_CHANNEL_ACCESS_TOKEN, PORT

# Command to run
CMD ["./server"]
