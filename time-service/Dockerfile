FROM golang:1.25-alpine AS builder
WORKDIR /app
COPY go.mod ./

COPY . .
RUN go build -o /go-app
FROM alpine:latest
WORKDIR /root/
COPY --from=builder /go-app .
EXPOSE 8081
CMD ["./go-app"]