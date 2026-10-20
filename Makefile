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

lint: ## runs the pinned golangci-lint version via go run
	go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.14.0 run ./...

lint-fix: ## automatically fix lint issues where possible
	golangci-lint run --fix ./...

check: fmt-check vet test lint

build:
	go build -o ai-issue ./cmd/ai-issue

release-check: ## validates .goreleaser.yaml without building anything
	goreleaser check

snapshot: ## builds all release targets locally without publishing, for pre-tag verification
	goreleaser release --snapshot --clean
