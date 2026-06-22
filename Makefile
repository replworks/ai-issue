fmt: ## runs go formatter
	go fmt ./...

fmt-check:
	@files=$$(gofmt -l .); \
	if [ -n "$$files" ]; then \
		echo "Files not formatted:"; \
		echo "$$files"; \
		exit 1; \
	fi

test:
	go test ./...

vet:
	go vet ./...

lint: ## runs golangci-lint via go run
	golangci-lint run ./...

lint-fix: ## automatically fix lint issues where possible
	golangci-lint run --fix ./...

check: fmt-check vet test

build:
	go build -o ai-issue ./cmd/ai-issue
