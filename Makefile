.PHONY: build-app

# Build commands
build-app:
	go build  -o ./bin/app ./cmd/app/main.go

build-migrator:
	go build  -o ./bin/migrator ./cmd/migrator/main.go

# Run commands
run-app-local:
	docker-compose -f docker-compose.local.yml up &
	CONFIG_PATH=./config/local.yaml ./bin/app

run-app-prod:
	CONFIG_PATH=./config/prod.yaml ./bin/app

run-migrator-local:
	CONFIG_PATH=./config/local.yaml ./bin/migrator

run-migrator-prod:
	CONFIG_PATH=./config/prod.yaml ./bin/migrator


.DEFAULT_GOAL := build-app