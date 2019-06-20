TAG = v$(shell date +%Y%m%d.%H%M%S)

build:
	docker build --build-arg version=$(TAG) -t skaji/kubernetes-playground-pod:v$(TAG) .

push:
	docker push skaji/kubernetes-playground-pod:v$(TAG)
