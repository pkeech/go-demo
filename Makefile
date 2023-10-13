build:
	@go build -o bin/godemo

run: build
	@./bin/godemo

test:
	@go test -v ./...