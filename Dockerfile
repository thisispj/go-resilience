# Build stage
FROM golang:1.23-alpine AS builder

LABEL authors="akshay.panikulamjoy"

# Set working directory
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download dependencies
RUN go mod download

# Copy source code
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o resilient-service ./cmd/api

# Final stage
FROM alpine:3.16

# Install ca-certificates
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy binary from builder
COPY --from=builder /app/resilient-service .

# Expose port
EXPOSE 8080

# Run the binary
CMD ["./resilient-service"]