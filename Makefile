# Target to build the Docker image
.PHONY: up
up:
	docker compose pull
	docker compose up -d


# Target to build the Docker image
.PHONY: down
down:
	docker compose down