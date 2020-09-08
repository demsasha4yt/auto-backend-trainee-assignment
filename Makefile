.PHONY: build
build:
	go build -v ./cmd/auto

.PHONY: test
test:
	go test -v -race ./...


.DEFAULT_GOAL := build

# migrate create -ext sql -dir migrations -seq create_logs_users_table
# migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations up
# migrate -database "postgres://postgres:pass@db/data?sslmode=disable" -path migrations down
