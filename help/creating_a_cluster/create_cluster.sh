#!/usr/bin/env bash

set -o xtrace

cd /creator
gcloud auth activate-service-account --key-file gcp-key.json
gcloud container clusters create --project $1 --zone $2 $3
kubectl create -f helm-service-account.yml
helm init --service-account helm
