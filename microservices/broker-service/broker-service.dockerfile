FROM golang:1.20.3-alpine3.17 as builder

RUN mkdir /app

COPY . /app

WORKDIR /app


RUN CGO_ENABLED=0 go build -o brokerApp  ./cmd/api/*.go

RUN chmod +x /app/brokerApp



FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/brokerApp /app


CMD ["/app/brokerApp"]