# dockerfile
# Сборка и запуск на Debian (glibc). Линкуемся с vendored librdkafka.
FROM golang:1.25-bookworm AS builder

WORKDIR /app
RUN apt-get update && apt-get install -y --no-install-recommends \
    build-essential pkg-config ca-certificates && rm -rf /var/lib/apt/lists/*

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags "-s -w" -o /out/service ./cmd/collector

FROM debian:bookworm-slim
RUN apt-get update && apt-get install -y --no-install-recommends \
    ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /out/service ./service

EXPOSE 4343
CMD ["./service"]

