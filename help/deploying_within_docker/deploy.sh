#!/usr/bin/env bash

set -o xtrace

cd /deployer
gcloud auth activate-service-account --key-file gcp-key.json
gcloud auth configure-docker
cd /app
web-deployer deploy --verbose $1 $2
