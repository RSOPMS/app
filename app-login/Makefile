.PHONY: all
all: test vet build

.PHONY: dev
dev:
	@air

.PHONY: build
build:
	@go build -o ./bin/service ./cmd/main.go

.PHONY: test
test:
	@go test ./...

.PHONY: vet
vet:
	@go vet ./...
