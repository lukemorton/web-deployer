build:
	go get ./...
	go install github.com/lukemorton/web-deployer/cmd/web-deployer

test: build
	go test ./...
	# exec $(GOPATH)/bin/web-deployer

test_publish: build
	cd internal/fixtures/ruby && exec $(GOPATH)/bin/web-deployer publish staging $(VERSION)

test_deploy: build
	cd internal/fixtures/ruby && exec $(GOPATH)/bin/web-deployer deploy staging $(VERSION)

charts/dist:
	mkdir -p charts/dist

publish_charts: charts/dist
	helm package charts/web-app -d charts/dist
	helm repo index --url https://web-deployer-charts.storage.googleapis.com charts/dist
	gsutil rsync -d charts/dist gs://web-deployer-charts
