FROM golang:latest as base

FROM base as protoc-compiler
WORKDIR /app

COPY .. /app

RUN apt-get update \
 && DEBIAN_FRONTEND=noninteractive \
    apt-get install --no-install-recommends --assume-yes \
      protobuf-compiler \
      libprotobuf-dev

RUN go mod tidy
RUN go get -u github.com/golang/protobuf/protoc-gen-go

RUN cd /go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 && make build

#CMD go run consume.go

FROM base as prod

WORKDIR /app

COPY .. /app

RUN go mod tidy

CMD go run main.go publish