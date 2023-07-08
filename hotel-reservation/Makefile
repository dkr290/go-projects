build:  
	@go build -o bin/api

run: build
	@./bin/api

test:
	@go test -v ./...

mongodb:
	@docker run -d  --name mongodb  -v ${PWD}/db-data:/data/db -p 27017:27017  mongo:latest 


mongodb_start:
	@docker start mongodb



stop:
	
	@docker stop mongodb
