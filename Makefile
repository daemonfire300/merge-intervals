lint:
	go run -mod=vendor github.com/golangci/golangci-lint/cmd/golangci-lint run
test:
	go test -v -count=1 -mod=vendor -short ./...
docker-build:
	docker build -t mergectl .
build:
	go build -v -o mergectl ./cmd/mergectl