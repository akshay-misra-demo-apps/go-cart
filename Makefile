# Target to build the Docker image
.PHONY: up
up:
	docker-compose up --pull

# Target to build the Docker image
.PHONY: down
down:
	docker-compose down