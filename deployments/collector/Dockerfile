# dockerfile
# Этап сборки на Alpine (musl) для совместимости с финальным образом
FROM golang:1.23-alpine AS builder

WORKDIR /app

# Инструменты для сборки cgo и dev-зависимости
RUN apk add --no-cache build-base pkgconf librdkafka-dev

# Кешируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники и собираем
COPY . ./
RUN CGO_ENABLED=1 GOOS=linux go build -trimpath -ldflags "-s -w" -o /out/service ./cmd/collector

# Финальный образ
FROM alpine:3.19

# Рантайм-зависимости (без *-dev)
RUN apk add --no-cache tzdata ca-certificates librdkafka

WORKDIR /app
COPY --from=builder /out/service ./service

EXPOSE 4343
CMD ["./service"]
