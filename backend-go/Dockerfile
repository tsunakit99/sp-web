FROM golang:1.24.2 AS builder

WORKDIR /app
COPY . .
RUN go mod download
RUN go build -o app ./cmd/main.go

FROM alpine:3.20
RUN apk add --no-cache ca-certificates

WORKDIR /app
COPY --from=builder /app/app .

CMD ["./app"]
