.PHONY: download run test build

PROJECT_NAME="tabnews-go"
TEST_PKGS=./pkg/... ./cmd/...

download:
	go mod download
	go mod tidy

run:
	go run cmd/$(PROJECT_NAME)/main.go

test:
	go test -v $(TEST_PKGS)

coverage:
	go test -coverprofile=coverage.out $(TEST_PKGS)
	go tool cover -html=coverage.out

build: download
	go build ./cmd/$(PROJECT_NAME)

infra-up:
	docker-compose -f infra/compose.yaml up -d

infra-down:
	docker-compose -f infra/compose.yaml down
