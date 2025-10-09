# ==== Build stage ====
FROM golang:1.25-alpine AS builder

WORKDIR /app

RUN apk add --no-cache gcc musl-dev

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN cp .env.example .env

ENV CGO_ENABLED=1

RUN go build -o server ./src/main.go

# ==== Runtime stage ====
FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache sqlite-libs

COPY --from=builder /app/server .
COPY --from=builder /app/.env .env
COPY --from=builder /app/data ./data

RUN mkdir -p db

EXPOSE 3000

CMD ["./server"]
