.SILENT:

BINARY_NAME=nasa-neows-cli-tool
DOCKER_IMAGE=nasa-neows-cli-tool
DOCKER_CONTAINER=nasa-neows-cli-tool

build:
	docker build --tag $(DOCKER_IMAGE) .
	docker run --name $(DOCKER_IMAGE) -e API_KEY=$(API_KEY) $(DOCKER_CONTAINER)

run:
	docker start -a $(DOCKER_CONTAINER)

test:
	API_KEY=$(API_KEY) go test ./...

remove:
	docker rm $(DOCKER_CONTAINER)
	docker rmi $(DOCKER_IMAGE)

run-dev:
	API_KEY=$(API_KEY) go run ./main.go

build-binary:
	go build -o $(BINARY_NAME) ./main.go

run-binary:
	API_KEY=$(API_KEY) ./$(BINARY_NAME)

remove-binary:
	go clean $(BINARY_NAME)