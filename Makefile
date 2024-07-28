.PHONY: build-app
build-app:
	go build  -o ./bin/app ./cmd/app/main.go

build-migrator:
	go build  -o ./bin/migrator ./cmd/migrator/main.go

run-local:
	docker-compose -f docker-compose.local.yml up &
	CONFIG_PATH=./config/local.yaml ./bin/app

run-migrator:
	CONFIG_PATH=./config/local.yaml ./bin/migrator


.DEFAULT_GOAL := build-app