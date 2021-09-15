
# VARIABLES
# -


# CONFIG
.PHONY: help print-variables
.DEFAULT_GOAL := help


# ACTIONS

## build

__check-container-tag :
	@[ "$(CONTAINER_TAG)" ] || ( echo "Missing container tag (CONTAINER_TAG), please define it and retry"; exit 1 )

build-all :		## Build all applications
	cd producer-random && go mod tidy && go build ./...
	cd producer-api && go mod tidy && go build ./...
	cd consumer-logger && go mod tidy && go build ./...
	cd consumer-mongo-db && go mod tidy && go build ./...

build-all-containers : __check-container-tag		## Build all containers
	cd producer-random && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	#cd producer-api && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	cd consumer-logger && make container-build CONTAINER_TAG=$(CONTAINER_TAG)
	#cd consumer-mongo-db && make container-build CONTAINER_TAG=$(CONTAINER_TAG)

push-all-containers : __check-container-tag		## Push all containers to Docker Hub
	cd producer-random && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	#cd producer-api && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	cd consumer-logger && make container-push CONTAINER_TAG=$(CONTAINER_TAG)
	#cd consumer-mongo-db && make container-push CONTAINER_TAG=$(CONTAINER_TAG)

build-push-all-containers : build-all-containers push-all-containers		## Build and push all containers to Docker Hub

## run

start-doc-comp :		## Run examples in Docker compose
	-@docker swarm init
	docker stack deploy kafka -c docker-compose.yaml

stop-doc-comp :		## Run examples in Docker compose
	-@docker stack rm kafka
	sleep 30
	docker swarm leave --force


## helpers

help :		## Help
	@echo ""
	@echo "*** \033[33mMakefile help\033[0m ***"
	@echo ""
	@echo "Targets list:"
	@grep -E '^[a-zA-Z_-]+ :.*?## .*$$' $(MAKEFILE_LIST) | sort -k 1,1 | awk 'BEGIN {FS = ":.*?## "}; {printf "\t\033[36m%-30s\033[0m %s\n", $$1, $$2}'
	@echo ""

print-variables :		## Print variables values
	@echo ""
	@echo "*** \033[33mMakefile variables\033[0m ***"
	@echo ""
	@echo "- - - makefile - - -"
	@echo "MAKE: $(MAKE)"
	@echo "MAKEFILES: $(MAKEFILES)"
	@echo "MAKEFILE_LIST: $(MAKEFILE_LIST)"
	@echo "- - -"
	@echo ""
