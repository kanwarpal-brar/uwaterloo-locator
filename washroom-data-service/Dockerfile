# Build stage
FROM golang:1.23.6-alpine AS builder

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Set working directory
WORKDIR /app

# Copy go mod files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=1 GOOS=linux go build -a -ldflags="-w -s" -o washroom-service .

# Production stage
FROM alpine:latest

# Install runtime dependencies
RUN apk add --no-cache sqlite-libs ca-certificates

# Create non-root user
RUN adduser -D appuser

# Create app directory and set permissions
WORKDIR /app
RUN chown appuser:appuser /app

# Copy binary from builder
COPY --from=builder /app/washroom-service .

# Copy SQLite database file if it exists
COPY --from=builder /app/washrooms.db ./washrooms.db
RUN chown appuser:appuser washrooms.db

# Switch to non-root user
USER appuser

# Expose port
EXPOSE 8080

# Set environment variables
ENV GIN_MODE=release

# Run the application
CMD ["./washroom-service"]
