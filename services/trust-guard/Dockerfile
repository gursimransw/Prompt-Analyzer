# Build stage
FROM golang:1.25.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build \
    -o trust-guard \
    ./cmd/server

# Runtime stage
FROM alpine:3.21

WORKDIR /app

RUN addgroup -S bearbreach && adduser -S bearbreach -G bearbreach

COPY --from=builder /app/trust-guard ./trust-guard
COPY config ./config

USER bearbreach

EXPOSE 8000

ENTRYPOINT ["./trust-guard"]
CMD ["-config", "config/server/config.yaml"]