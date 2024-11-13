all: test vet build

dev:
	@air

build:
	@go build -o ./bin/service ./cmd/main.go

test:
	@go test ./...

vet:
	@go vet ./...