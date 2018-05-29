# Deploying within Docker

Deploying from within the `lukemorton/web-deployer-tools` Docker image is simple.

We've put together a `Makefile` that authenticates and then deploys to Google Kubernetes Engine.

## Usage

```
make DEPLOYMENT=staging VERSION=v1.2.3
```

This will then log you into Google Cloud with `gcloud auth activate-service-account` before setting up a new cluster and installing help within it with a secure service account.

You need to make sure a `gcp-key.json` file exists in this directory container  credentials for `gcloud auth activate-service-account` to work.
