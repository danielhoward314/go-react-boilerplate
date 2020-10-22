.DEFAULT_GOAL := help
SHA := latest
TAG := another-test

## Change next line to push to your Docker registry
IMAGE_TAG := registry.hub.docker.com/repositories/danielhoward314/go-react-boilerplatedanielhoward314/go-react-boilerplate

build: ## Build the docker container and tag as latest
	docker build -t ${IMAGE_TAG}:${TAG} .
.PHONY: build

shell: build ## Build the docker container and then run in interaction mode
	docker run -it ${IMAGE_TAG}:${TAG} /bin/sh
.PHONY: shell

push: ## Push the docker container to registry
	docker push ${IMAGE_TAG}:${TAG}
.PHONY: push

tag: ## Tag the docker image
	docker tag ${IMAGE_TAG}:${SHA} ${IMAGE_TAG}:${TAG}
.PHONY: tag

help: ## show this usage
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n\nTargets:\n"} /^[a-zA-Z_-]+:.*?##/ { printf "  \033[36m%-10s\033[0m %s\n", $$1, $$2 }' $(MAKEFILE_LIST)
.PHONY: help
