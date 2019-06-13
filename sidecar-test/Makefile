VERSION = 0.0.1
COMMITISH = $(shell git rev-parse --short HEAD)
IMAGE_NAME_PREFIX = skaji/sidecar-test

build:
	cd origin && docker build --build-arg port=9000 --build-arg version=$(VERSION) --build-arg commitish=$(COMMITISH) -t $(IMAGE_NAME_PREFIX)-origin .
	cd proxy  && docker build --build-arg port=8080 --build-arg version=$(VERSION) --build-arg commitish=$(COMMITISH) -t $(IMAGE_NAME_PREFIX)-proxy  .
