.PHONY: build-app
build-app:
	go build  -o ./bin/app ./cmd/app/main.go

run-local:
	CONFIG_PATH="./config/local.yaml" ./bin/app

.DEFAULT_GOAL := build-app