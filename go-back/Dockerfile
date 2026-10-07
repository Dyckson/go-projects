# -------------------------
# Build stage
# -------------------------
FROM golang:1.27.1-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# -------------------------
# Runtime stage
# -------------------------
FROM alpine:3.24.2

WORKDIR /app

RUN adduser -D appuser

COPY --from=builder /app/main .

USER appuser

EXPOSE 1111

CMD ["./main"]