FROM golang:alpine AS builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o gate-backend ./cmd/api/main.go

FROM alpine:latest
WORKDIR /app

COPY --from=builder /app/gate-backend .

EXPOSE 3000

CMD ["./gate-backend"]