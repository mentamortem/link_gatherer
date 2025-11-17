# dockerfile
# Сборка и запуск на Debian (glibc), используем vendored librdkafka от confluent-kafka-go
FROM golang:1.25-bookworm AS builder

WORKDIR /app

# Инструменты для cgo
RUN apt-get update && apt-get install -y --no-install-recommends build-essential pkg-config ca-certificates && rm -rf /var/lib/apt/lists/*

# Кеш зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Сборка (без -tags dynamic — линкуемся с vendored glibc librdkafka)
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags "-s -w" -o /out/service ./cmd/collector

# Финальный образ на Debian
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates tzdata && rm -rf /var/lib/apt/lists/*

WORKDIR /app
COPY --from=builder /out/service ./service

EXPOSE 4343
CMD ["./service"]
