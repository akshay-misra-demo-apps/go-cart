# Define variables
DOCKER_HARBOR = docker.io
DOCKER_USER = akshaymisra24
DOCKER_REPO = go-cart-product-management
DOCKER_IMAGE_TAG = 0.1.0

# Target to build the Docker image
build:
	docker build -t $(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG) .

# Target to push the Docker image to Docker Harbor
push:
	echo "DOCKER_PASSWORD: $(DOCKER_PASSWORD)"
	echo "$(DOCKER_PASSWORD)" | docker login -u $(DOCKER_USER) --password-stdin $(DOCKER_HARBOR)
	docker push $(DOCKER_USER)/$(DOCKER_REPO):$(DOCKER_IMAGE_TAG)

# Target to build and push the Docker image
build_and_push: build push
