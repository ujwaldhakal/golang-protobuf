FROM golang:latest

WORKDIR /app

COPY ./ /app


RUN go mod tidy

CMD go run main.go publish