
run:
	go run ./src/.

build:
	go build -o ./bin/service ./src/.

fmt:
	go fmt ./src/...

dep:
	@ cd src
	go mod tidy

.PHONY: run
