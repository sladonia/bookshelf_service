# bookshelf_service

A simple http REST API service made in self-education purpose. Uses postgres as a db backend. No ORM is used.

### build dependencies

* go v1.13.5
* GNU Make
* docker v19.03.3
* docker-compose v1.25.0
* psql

### build

##### Development environment

build executable
```sh
make build
```

build bookshelf service docker image
```sh
docker-compose build bookshelf
```

up postgres and postgres_test services
```shell script
docker-compose up -d postgres postgres_test
```

create tables in postgres
```sh
make create_tables
```

run tests
```sh
make test
```

up bookshelf service
```sh
docker-compose up bookshelf
```

check if service alive
```sh
curl localhost:8080
```
response should be:
```json
{"message":"Welcome to bookshelf_service api"}
```

##### Production environment

build bookshelf service docker image
```sh
make docker_build
```
 remember to create database & tables
 
### configuration

service is configurable with env vars and the defaults are

```.env
SERVICE_NAME=bookshelf_service
ENV=dev
LOG_LEVEL=debug
PORT=:8080

POSTGRES_HOST=localhost
POSTGRES_PORT=5432
POSTGRES_DB=bookshelf_db
POSTGRES_USER=user
POSTGRES_PASSWORD=password
POSTGRES_MAX_OPEN_CONNECTIONS=25
POSTGRES_MAX_IDLE_CONNECTIONS=25
POSTGRES_CONNECTION_MAX_LIFETIME=5
``` 