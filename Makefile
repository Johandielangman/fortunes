run: build
	@./bin/fortune | cowsay

build:
	@go build -o ./bin/fortune fortune.go
