lint:
	go run -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint run
test:
	go test -v -count=1 -mod=vendor -short ./...