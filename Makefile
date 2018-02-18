build:
	go get ./...
	go install github.com/lukemorton/web-deployer/cmd/web-deployer

test: build
	exec $(GOPATH)/bin/web-deployer
	exec $(GOPATH)/bin/web-deployer publish --k8s-project doorman-1200 internal/fixtures/ruby

charts/dist:
	mkdir -p charts/dist

publish_charts: charts/dist
	helm package charts/web-app -d charts/dist
	helm repo index --url https://web-deployer-charts.storage.googleapis.com charts/dist
	gsutil rsync -d charts/dist gs://web-deployer-charts
