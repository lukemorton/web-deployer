all: gcloud/login gcloud/create-cluster kubectl/helm-service-account helm/init

gcloud/login:
	gcloud auth login

gcloud/create-cluster:
	gcloud container clusters create --project $(PROJECT) --zone $(ZONE) $(CLUSTER)

kubectl/helm-service-account:
	kubectl create -f helm-service-account.yml

helm/init:
	helm init --service-account helm
