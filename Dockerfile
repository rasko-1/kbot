FROM golang:tip-alpine3.21 AS builder
RUN apk add --no-cache ca-certificates
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download

COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o kbot ./main.go

FROM scratch
WORKDIR /
COPY --from=builder /app/kbot .
COPY --from=builder /app/.env .
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
ENTRYPOINT ["/kbot"]
