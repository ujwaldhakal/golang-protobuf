.PHONY: *

start:
	docker-compose up -d
publish:
	docker-compose run app go run publish.go

consume:
	docker-compose run app go run consume.go

format:
	gofmt -s -w .

build-proto:
	 docker-compose run protoc-compiler protoc \
	  -I . \
	  -I /go/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
	  --go_out=":messages" \
	  --validate_out="lang=go:messages" \
	  	messages/*.proto


asd:
	 docker-compose run protoc-compiler