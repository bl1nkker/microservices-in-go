# base go image
FROM golang:1.18-alpine as builder

RUN mkdir /app

WORKDIR /app

COPY . /app

RUN CGO_ENABLED=0 go build -o authenticationApp ./cmd/api

RUN chmod +x /app/authenticationApp


FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/authenticationApp /app

CMD ["/app/authenticationApp"]