SERVICE_NAME=bookshelf_service
BIN=app
BUILD=CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./bin/$(BIN) ./src/.

run:
	@ go run ./src/.

build:
	$(BUILD)

fmt:
	go fmt ./src/...

dep:
	@ cd src
	go mod tidy

update:
	@ cd src
	go get -u

docker-build:
	$(BUILD)
	docker build -t $(SERVICE_NAME) .

stop-containers:
	docker container ls -q | xargs docker container stop 2>/dev/null || true

docker-run:
	@ docker run -p 8080:8080 bookshelf_service

.PHONY: run build fmt dep update docker-build stop-containers
