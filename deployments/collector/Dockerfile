# Этап сборки
FROM golang:1.25 AS builder

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для кеширования зависимостей
COPY go.mod ./
RUN go mod download

# Копируем исходники
COPY cmd/ cmd/
COPY internal/ internal/

# Собираем бинарь
RUN CGO_ENABLED=0 GOOS=linux go build -o service ./cmd/link_collector

# Финальный образ
FROM alpine:3.19

# Устанавливаем часовой пояс и сертификаты (если сервис ходит по HTTPS)
RUN apk add --no-cache tzdata ca-certificates

RUN mkdir /app
WORKDIR /app

# Копируем бинарь из билдера
COPY --from=builder /app/service .

# Открываем порт (замени на свой)
EXPOSE 1111

# Запускаем сервис
CMD ["./service"]
