.PHONY: test
test:
	go test -v -race -cover -coverprofile coverage.out -count 1 .

lint:
	golangci-lint run .
