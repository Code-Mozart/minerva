build:
	@go build -C ./src -o ../bin/minerva-backend ./cmd

run: build
	@./bin/minerva-backend

test:
	@go test -C src -v ./...
