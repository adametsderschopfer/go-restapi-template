.PHONY: build-app
build-app:
	go build  -o ./bin/app ./cmd/app/main.go

run-local:
	docker-compose -f docker-compose.local.yml up -d
	CONFIG_PATH="./config/local.yaml" ./bin/app

.DEFAULT_GOAL := build-app