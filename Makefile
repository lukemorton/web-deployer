build:
	go get ./...
	go install github.com/lukemorton/web-deployer/cmd/web-deployer

test: build
	go test ./...
	# exec $(GOPATH)/bin/web-deployer

test_publish_ruby: build
	cd internal/fixtures/ruby && exec $(GOPATH)/bin/web-deployer publish --verbose staging $(VERSION)

test_deploy_ruby: build
	cd internal/fixtures/ruby && exec $(GOPATH)/bin/web-deployer deploy --verbose staging $(VERSION)

test_publish_dotnet: build
	cd internal/fixtures/dotnet && exec $(GOPATH)/bin/web-deployer publish --verbose staging $(VERSION)

test_deploy_dotnet: build
	cd internal/fixtures/dotnet && exec $(GOPATH)/bin/web-deployer deploy --verbose staging $(VERSION)

charts/dist:
	mkdir -p charts/dist

publish_charts: charts/dist
	helm package charts/web-app -d charts/dist
	helm repo index --url https://web-deployer-charts.storage.googleapis.com charts/dist
	gsutil rsync -d charts/dist gs://web-deployer-charts

publish_images:
	cd images/tools && make publish
