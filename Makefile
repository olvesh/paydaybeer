DOCKER_PRIVATE_REPO ?= vimond-dockerv2-local.jfrog.io
NAME = $(DOCKER_PRIVATE_REPO)/paydaybeer
VERSION = $(shell git describe --tags --always)


.PHONY: build test tag_latest release ssh

compile:
	go get -t -d  ./...
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build


build: compile
	docker build  -t $(NAME):$(VERSION)  .

test: build
  go test
	docker run --rm $(NAME):$(VERSION)

tag_latest:
	docker tag  $(NAME):$(VERSION) $(NAME):latest

release: build
	@if ! docker images $(NAME) | awk '{ print $$2 }' | grep -q -F $(VERSION); then echo "$(NAME) version $(VERSION) is not yet built. Please run 'make build'"; false; fi
	docker push $(NAME)