
FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/todo-app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todo-app .
EXPOSE 8080
CMD ["./todo-app"]FROM golang:1.23-alpine AS builder
RUN apk add --no-cache git
WORKDIR /app
COPY . .
RUN go mod tidy && \
    CGO_ENABLED=0 GOOS=linux go build -o todo-app ./cmd/todo-app

FROM alpine:latest
WORKDIR /root/
COPY --from=builder /app/todo-app .
EXPOSE 8080
CMD ["./todo-app"]

