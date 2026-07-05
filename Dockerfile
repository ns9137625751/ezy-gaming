# ── Stage 1: Builder ─────────────────────────────────────────
FROM golang:1.23-alpine AS builder

RUN apk add --no-cache git ca-certificates

WORKDIR /build

COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o main ./cmd/server

# ── Stage 2: Runner ──────────────────────────────────────────
FROM alpine:3.20

RUN apk add --no-cache ca-certificates tzdata

WORKDIR /app

COPY --from=builder /build/main .
COPY --from=builder /build/templates ./templates
COPY --from=builder /build/static    ./static

RUN mkdir -p static/uploads

EXPOSE 8080

CMD ["./main"]
