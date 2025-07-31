run: build
	@ echo "starting server..."
	./bin/server

build:
	go build -o ./bin/server ./cmd/api/main.go

export DB_TEST_PATH=$(shell pwd)/data/test.db

test: migrate_test
	@ go test ./pkg/repositories/sqlite/../... -v

bench: migrate_test
	@ cd ./pkg/repositories/sqlite/student/ && go test -bench=.

migrate_test:
	@ GOOSE_DBSTRING=data/test.db goose down-to 0
	@ GOOSE_DBSTRING=data/test.db goose up

migrate_prod:
	@ goose down-to 0
	@ goose up
