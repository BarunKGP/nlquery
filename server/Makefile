build: 
	@go build -o bin/main 

run: build
	@./bin/main

docker-build:
	@go mod download
	CGO_ENABLED=0 GOOS=linux go build -o /nlquery
