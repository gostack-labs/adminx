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
	docker exec -it postgres12 createdb --username=root -owner=root adminx

dropdb:
	docker exec -it postgres12 dropdb adminx

lint:
	golangci-lint run

.PHONY: dbml2sql dbdocs postgres createdb dropdb lint

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

