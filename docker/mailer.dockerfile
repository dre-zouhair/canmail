# base go image
FROM golang:1.18-alpine as builder

WORKDIR /app

COPY ../mailer /app

RUN go build -o mailer ./cmd

RUN chmod +x /app/mailer

# build a tiny docker image
FROM alpine:latest

RUN mkdir /app

COPY --from=builder /app/mailer /app

CMD [ "/app/mailer" ]