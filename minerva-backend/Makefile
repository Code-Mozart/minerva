build:
	@go build -C ./src -o ../bin/minerva-backend ./cmd

run: build
	@./bin/minerva-backend

test:
	@go test -C src -v ./...

docker-build:
	@docker build -t minerva-backend:latest .

docker-run: docker-build
	@docker run -p 8134:8134 minerva-backend:latest
