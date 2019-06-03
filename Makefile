GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
BINARY_NAME=runner-auto
IMAGE_NAME=subaruqui/gitlab-runner
all: build build-docker
build:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
run:
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 $(GOBUILD) -o $(BINARY_NAME) -v
	./$(BINARY_NAME)
deps:
	 $(GOGET) github.com/google/uuid

build-docker:
	docker build -t $(IMAGE_NAME) .
