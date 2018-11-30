BIN_PATH = ../bin
SRC_PATH = ./app
DOCKER_RUN = docker-compose run --rm --user app golang

-include .env

.PHONY: vendor

# Run application
run:
	docker-compose run --rm --user app --publish $(APP_PORT):$(APP_PORT) golang go run main.go

# Make build
build:
	$(DOCKER_RUN) go build -o $(BIN_PATH)/$(APP_BIN) main.go

# Go get a package
get:
	$(DOCKER_RUN) dep ensure -add ${pkg}

# Update vendor
vendor:
	$(DOCKER_RUN) /bin/sh -c 'cd $(SRC_PATH); dep ensure'

# Enter into golang container
bash:
	$(DOCKER_RUN) bash

# Configure environment
install:
	if [ ! -f ".env" ]; then cp ./.env.example ./.env; fi
	docker-compose build --force-rm --no-cache
	$(MAKE) up
	$(MAKE) vendor

# Start docker
up: down
	docker-compose up --build -d

# Stop docker
down:
	docker-compose down

# Show docker processes
ps:
	docker-compose ps
