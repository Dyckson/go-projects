# -------------------------
# Build stage
# -------------------------
FROM golang:1.24.5 AS builder

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main .


# -------------------------
# Runtime stage
# -------------------------
FROM alpine:3.22

WORKDIR /app

COPY --from=builder /app/main .

EXPOSE 1111

CMD ["./main"]
