FROM golang:1.24.5 AS builder

# Копируем сначала только модули для кэширования
COPY go.mod go.sum ./
RUN go mod download

# Рабочая директория = корень проекта
WORKDIR /app

# Копируем весь проект
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o chetam ./cmd/app/.

FROM alpine:3.20

WORKDIR /usr/bin

COPY --from=builder /app/chetam /usr/bin/chetam

COPY entrypoint.sh /entrypoint.sh
RUN chmod +x /entrypoint.sh

ENTRYPOINT ["/entrypoint.sh"]