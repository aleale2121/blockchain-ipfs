# Build stage
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go build -o main cmd/main.go

# Run stage
FROM alpine:3.16

WORKDIR /app
COPY --from=builder /app/main .
COPY app.env .
COPY start.sh .
COPY assets /app/assets
RUN chmod +x /app/main

EXPOSE 9090

CMD ["/bin/sh", "/app/start.sh"]
