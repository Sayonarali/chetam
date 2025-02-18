# Официальный образ Go
FROM golang:1.23.2

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .
RUN go mod tidy

RUN go build -o /app/main ./cmd/app

# Открываем порт
EXPOSE 8080

CMD ["./main"]
