# Create a cluster on Google Kubernetes Engine

One of the easiest and most production ready ways to creating a Kubernetes cluster is to use GKE (Google Kubernetes Engine).

We've put together a `Makefile` that creates a brand new cluster to help with testing web-deployer. Feel free to use it too.

## Usage

```
make PROJECT=my-project ZONE=europe-west2-a CLUSTER=ruby-sample-app-cluster
```

This will then log you into Google Cloud with `gcloud auth activate-service-account` before setting up a new cluster and installing help within it with a secure service account.

You need to make sure a `gcp-key.json` file exists in this directory container  credentials for `gcloud auth activate-service-account` to work.
