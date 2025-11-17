# dockerfile
# Сборка на Alpine (musl) с динамическим линкованием librdkafka
FROM golang:1.25-alpine AS builder

WORKDIR /app

# Инструменты для cgo и dev-зависимости librdkafka
RUN apk add --no-cache build-base pkgconf librdkafka-dev openssl-dev zstd-dev cyrus-sasl-dev

# Кеш зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Сборка с тегом dynamic, чтобы использовать системную librdkafka (musl), а не vendored glibc
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -tags dynamic -trimpath -ldflags "-s -w" -o /out/service ./cmd/collector

# Финальный образ
FROM alpine:3.19

# Рантайм-зависимости librdkafka
RUN apk add --no-cache tzdata ca-certificates librdkafka openssl zstd-libs cyrus-sasl

WORKDIR /app
COPY --from=builder /out/service ./service

EXPOSE 4343
CMD ["./service"]
