# Etapa 1: build do binário Go
FROM golang:1.24-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o dist/yourbinary ./cmd/api

# Etapa 2: imagem enxuta só para rodar o binário
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/dist/yourbinary .

EXPOSE 8080
CMD ["./yourbinary"]
