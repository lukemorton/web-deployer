IMAGE := lukemorton/web-deployer-tools

all:
	docker run \
		--rm \
		-v `pwd`:/creator \
		-it \
		$(IMAGE) \
		bash /creator/create_cluster.sh $(PROJECT) $(ZONE) $(CLUSTER)
