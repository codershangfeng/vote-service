.PHONY: fmt
fmt:
	go fmt ./...

.PHONY: validate
validate:
	swagger validate ./api/swagger.yaml

.PHONY: gen
gen: validate
	rm -rf ./internal/api
	mkdir -p ./internal/api
	swagger generate server \
			--target=./internal/api \
			--spec=./api/swagger.yaml \
			--exclude-main \
			--name=vote-service

.PHONY: build
build: fmt validate gen
	mkdir -p ./bin
	CGO_ENABLED=0 GOOS=linux go build \
	-a \
	-installsuffix cgo \
	-o ./bin/vote-service ./cmd/main.go
	@echo "\033[0;32mSuccessfully build application in ./bin/vote-service\033[0m"

.PHONY: run
run: fmt
	go run ./cmd/main.go

.PHONY: install-swagger # Install go-swagger
install-swagger:
	@echo ">> Installing go-swagger"
	go get -u github.com/go-swagger/go-swagger/cmd/swagger

.PHONY: swagger-editor
swagger-editor:
	MY_VAR=$$(docker ps | grep swaggerapi/swagger-editor | cut -d" " -f 1) && \
    if [ -n "$$MY_VAR" ]; then \
	  echo "Stop exising swagger editor..."; \
      docker stop "$$MY_VAR"; \
    fi && \
	docker run -d --rm -p 80:8080 swaggerapi/swagger-editor &&\
    echo '$@: Successfully start Swagger editor'


# Test
.PHONY: clean-test-cache
clean-test-cache:
	go clean -testcache

.PHONY: utest
utest: fmt clean-test-cache
	go test -tags=unit ./...

.PHONY: itest
itest: fmt clean-test-cache
	go test -tags=integration ./...

.PHONY: test
test: fmt clean-test-cache gen 
	go test -tags="integration unit" ./...
