.PHONY: test lint run

test:
	go test -race ./...

lint:
	golangci-lint run

run:
	go run cmd/app/main.go -file test/url1000.txt
