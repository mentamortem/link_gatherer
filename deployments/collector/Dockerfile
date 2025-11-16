# Этап сборки
FROM golang:1.25 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходники
COPY . ./

# Собираем бинарь
RUN CGO_ENABLED=1 GOOS=linux go build -o service ./cmd/collector

# Финальный образ
FROM alpine:3.19

# Устанавливаем часовой пояс и сертификаты (если сервис ходит по HTTPS)
RUN apk add --no-cache tzdata ca-certificates pkgconf librdkafka-dev

RUN mkdir /app
WORKDIR /app

# Копируем бинарь из билдера
COPY --from=builder /app/service .

# Открываем порт (замени на свой)
EXPOSE 4343

# Запускаем сервис
CMD ["./service"]
