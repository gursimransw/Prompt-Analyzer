# Stage 1: Build the Go binary
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Copy dependency files first for better Docker layer caching
COPY go.mod go.sum ./
RUN go mod download

# Copy source code
COPY . .

# Build the application binary
RUN go build -o prompt-analyzer ./cmd/prompt-analyzer

# Stage 2: Runtime image
FROM alpine:latest

WORKDIR /app

# Copy only the compiled binary from builder stage
COPY --from=builder /app/prompt-analyzer .

# Copy config files required at runtime
COPY configs ./configs

EXPOSE 8000

CMD ["./prompt-analyzer"]