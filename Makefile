build:
	go get ./...
	go install github.com/lukemorton/web-deployer/cmd/web-deployer

test: build
	exec $(GOPATH)/bin/web-deployer
	cd internal/fixtures/ruby && exec $(GOPATH)/bin/web-deployer publish staging v1

charts/dist:
	mkdir -p charts/dist

publish_charts: charts/dist
	helm package charts/web-app -d charts/dist
	helm repo index --url https://web-deployer-charts.storage.googleapis.com charts/dist
	gsutil rsync -d charts/dist gs://web-deployer-charts
