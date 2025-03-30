FROM golang:1.22.3 AS builder

WORKDIR /app

COPY go.mod go.sum ./
COPY ./app /app
RUN go mod download

COPY . .

RUN go build -o main ./cmd/main.go

# Imagen final
FROM debian:bullseye-slim

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 8084

CMD ["app/main"]