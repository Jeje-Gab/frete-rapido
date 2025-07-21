# Etapa 1: build do binário Go
FROM golang:1.22-alpine AS builder

WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN go build -o dist/yourbinary ./cmd

# Etapa 2: imagem enxuta só para rodar o binário
FROM alpine:latest

WORKDIR /app
COPY --from=builder /app/dist/yourbinary .
COPY ./env/.env.development .env    # Ajuste esse caminho se necessário

EXPOSE 8080
CMD ["./yourbinary"]
