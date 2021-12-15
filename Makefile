.PHONY: *

publish:
	docker-compose run app go run main.go publish

consume:
	docker-compose run app go run main.go consume

build-proto:
	 protoc \
	  -I . \
	  -I ${GOPATH}/pkg/mod/github.com/envoyproxy/protoc-gen-validate@v0.6.2 \
	  --go_out=":message" \
	  --validate_out="lang=go:message" \
	  	message/*.proto
