.PHONY: build
build:
	go build -v ./cmd/auto

.PHONY: test
test:
	go test -v -race ./...

.PHONY: migrate_up
migrate_up:
	migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations up

.PHONY: run
run: migrate_up build
	./auto

.DEFAULT_GOAL := build
