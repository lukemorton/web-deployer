charts/dist:
	mkdir -p charts/dist

publish_charts: charts/dist
	helm package charts/web-app -d charts/dist
	helm repo index --url https://web-deployer-charts.storage.googleapis.com charts/dist
	gsutil rsync -d charts/dist gs://web-deployer-charts
