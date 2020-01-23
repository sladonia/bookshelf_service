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

docker_build:
	$(BUILD)
	docker build -t $(SERVICE_NAME) .

stop_containers:
	docker container ls -q | xargs docker container stop 2>/dev/null || true

docker_run:
	@ docker run -p 8080:8080 $(SERVICE_NAME)

create_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/crete_tables.sql bookshelf_db user

recreate_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/recrete_tables.sql bookshelf_db user

clear_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/delete_all_data.sql bookshelf_db user

.PHONY: run build fmt dep update docker_build stop_containers docker_run create_tables
