# Define variables
DOCKER_HARBOR = docker.io
DOCKER_USER = akshaymisra24
DOCKER_REPO = go-cart-product-management
DOCKER_IMAGE_TAG = 0.1.0

# Target to build golang code
.PHONY: build
build:
	go build

# Target to run golang code
.PHONY: run
run:
	go run main.go

# Target to run tests and generate coverage report
.PHONY: test
test:
	go test -v -coverprofile coverage.out ./...
	go tool cover -html coverage.out -o coverage.html
	rm -f coverage.out

# Target to build the Docker image
.PHONY: docker_build
docker_build:
	docker build -t $(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG) .

# Target to push the Docker image to Docker Harbor
.PHONY: docker_push
docker_push:
	echo "$(DOCKER_PASSWORD)" | docker login -u $(DOCKER_USER) --password-stdin $(DOCKER_HARBOR)
	docker push $(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG)

# Target to build and push the Docker image
.PHONY: docker_build_push
docker_build_push: docker_build docker_push

# Target to run the Docker image
.PHONY: docker_run
docker_run:
	docker run -d -p 4100:4100 $(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG)

# Target to build and run the Docker container
.PHONY: docker_build_run
docker_build_run: docker_build docker_run

# Target to stop the Docker container
.PHONY: docker_stop
docker_stop:
	docker stop $$(docker ps -q --filter ancestor=$(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG))
