FROM golang:1.23.6-alpine AS builder

WORKDIR /app

# Install build dependencies
RUN apk add --no-cache gcc musl-dev sqlite-dev

# Copy go mod files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build with optimizations
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags="-w -s" -o washroom-service .

FROM alpine:3.19

WORKDIR /app

# Install runtime dependencies
RUN apk add --no-cache sqlite-libs ca-certificates tzdata

# Copy binary from builder
COPY --from=builder /app/washroom-service .

EXPOSE 8080

ENV GIN_MODE=release \
    ENVIRONMENT=production

CMD ["./washroom-service"]
