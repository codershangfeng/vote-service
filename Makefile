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
	-o ./bin/vote-service ./cmd/main.go
	@echo "\033[0;32mSuccessfully build application in ./bin/vote-service\033[0m"

.PHONY: run
run: fmt
	go run ./cmd/main.go

.PHONY: install-swagger # Install go-swagger
install-swagger:
	@echo ">> Installing go-swagger"
	go get -u github.com/go-swagger/go-swagger/cmd/swagger