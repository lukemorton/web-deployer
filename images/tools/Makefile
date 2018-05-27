IMAGE := lukemorton/web-deployer-tools

dist/web-deployer:
	CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o dist/web-deployer github.com/lukemorton/web-deployer/cmd/web-deployer

build: dist/web-deployer
	docker build -t $(IMAGE) .

run: build
	docker run --rm -t -i $(IMAGE) ash

publish: build
	docker push $(IMAGE)

clean:
	rm -rf dist
