export REGISTRY_NAME?=docker.io/arainho
export VERSION=1.0.0

.PHONY: build dockerhub-image

build: dockerhub-image

dockerhub-image: export PROJECT_NAME=api-sec-toobox
dockerhub-image:
	docker build --label version=$(VERSION) --build-arg TOOLBOX_VERSION=$(VERSION) --no-cache -t $(REGISTRY_NAME)/$(PROJECT_NAME):$(VERSION) .
