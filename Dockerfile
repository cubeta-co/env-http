FROM golang:1.16-alpine AS builder
WORKDIR /app
COPY go.* ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -o env-http main.go

FROM scratch

EXPOSE 8080
WORKDIR /app

COPY --from=builder /app/env-http /app/env-http

CMD ["./env-http"]
