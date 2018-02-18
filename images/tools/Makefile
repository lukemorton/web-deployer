IMAGE := lukemorton/web-deployer-tools

build:
	docker build -t $(IMAGE) .

run: build
	docker run --rm -t -i $(IMAGE) ash

publish: build
	docker push $(IMAGE)
