IMAGE := lukemorton/web-deployer-tools

all:
	docker run \
		--rm \
		-v /var/run/docker.sock:/var/run/docker.sock \
		-v `pwd`/../../internal/fixtures/ruby:/app \
		-v `pwd`:/deployer \
		-it \
		$(IMAGE) \
		bash /deployer/deploy.sh $(DEPLOYMENT) $(VERSION)
