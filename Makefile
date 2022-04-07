DB_URL=postgresql://root:secret@localhost:5432/adminx?sslmode=disable

# dbml to sql file
dbml2sql:
	dbml2sql --postgres -o docs/schema.sql docs/db.dbml

# generate online database documents
dbdocs:
	dbdocs build docs/db.dbml

# docker run postgresql database
postgres:
	docker run --name postgres12 -p 5432:5432 -e POSTGRES_USER=root -e POSTGRES_PASSWORD=secret -d postgres:12-alpine

# create database
createdb:
	docker exec -it postgres12 createdb --username=root --owner=root adminx

# drop database
dropdb:
	docker exec -it postgres12 dropdb adminx

args = `arg="$(filter-out $@,$(MAKECMDGOALS))" && echo $${arg:-${1}}`

# generate migrate file
migrate:
	migrate create -ext sql -dir internal/repository/db/migration --seq $(call args,init_schema)

# migrate up
migrateup:
	migrate -path internal/repository/db/migration --database "${DB_URL}" -verbose up

# migrate down
migratedown:
	migrate -path internal/repository/db/migration --database "${DB_URL}" -verbose down

# generate crud go code
sqlc:
	sqlc generate

# exec test
test:
	go test -cover ./...

# start server
server:
	go run main.go

# code specification inspection
lint:
	golangci-lint run

.PHONY: dbml2sql dbdocs postgres createdb dropdb migrate migrateup migratedown sqlc test server lint

# show help
help:
	@echo ''
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help

