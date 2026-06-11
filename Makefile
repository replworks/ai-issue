fmt:
	gofmt -w .

lint:
	golangci-lint run

test:
	go test ./...

check: fmt lint test
