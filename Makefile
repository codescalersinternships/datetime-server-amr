build http:
	go build cmd/http/main.go

build gin:
	go build cmd/gin/main.go

format:
	go fmt ./...

lint:
	golangci-lint run

build images:
	docker-compose build

launch containers:
	docker-compose up
