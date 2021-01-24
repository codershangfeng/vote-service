.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: validate
validate:
	swagger validate ./api/swagger.yaml

.PHONY: gen
gen: validate
	mkdir -p ./internal/api
	swagger generate server \
			--target=./internal/api \
			--spec=./api/swagger.yaml \
			--exclude-main \
			--name=vote-service

.PHONY: build
build: fmt validate gen
	mkdir -p ./bin
	go build \
	-a \
	-installsuffix cgo \
	-o ./bin ./cmd/main.go
	@echo "\033[0;32mSuccessfully build application in ./bin/main\033[0m"