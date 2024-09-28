build: 
	@go build -o bin/main 

run: build
	@./bin/main

docker-build:
	@go mod download
	CGO_ENABLED=0 GOOS=linux go build -o /nlquery
	docker compose up -d db --build
	goose postgres "postgres://nlquery:postgres@localhost:5432/nlquery" -dir "server/sql/schema" up

	
