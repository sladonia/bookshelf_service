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

test:
	@ go test -v -cover ./src/...

create_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/crete_tables.sql bookshelf_db user
	PGPASSWORD=password psql -h localhost -p 5433 -f postgres/crete_tables.sql bookshelf_db user

recreate_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/recrete_tables.sql bookshelf_db user
	PGPASSWORD=password psql -h localhost -p 5433 -f postgres/recrete_tables.sql bookshelf_db user

clear_tables:
	PGPASSWORD=password psql -h localhost -p 5432 -f postgres/delete_all_data.sql bookshelf_db user
	PGPASSWORD=password psql -h localhost -p 5433 -f postgres/delete_all_data.sql bookshelf_db user

.PHONY: run build fmt dep update docker_build test stop_containers docker_run create_tables
