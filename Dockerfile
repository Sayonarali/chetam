# Официальный образ Go
FROM golang:1.23.2

# Рабочая директория внутри контейнера
WORKDIR /app

# Копируем файлы проекта в контейнер
COPY . .
RUN go mod tidy

WORKDIR /app/cmd
RUN go build -o main

# Открываем порт
EXPOSE 8080

CMD ["./main"]
