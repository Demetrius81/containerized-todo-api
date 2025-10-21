FROM golang:1.25-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY go.mod go.sum ./
COPY vendor/ ./vendor/
COPY . .
RUN go build -mod=vendor -o /bin/server ./cmd/server

FROM alpine:latest
WORKDIR /app
COPY --from=builder /bin/server /app/server
EXPOSE 5002
CMD ["/app/server"]
