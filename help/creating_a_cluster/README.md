# Create a cluster on Google Kubernetes Engine

One of the easiest and most production ready ways to creating a Kubernetes cluster is to use GKE (Google Kubernetes Engine).

We've put together a `Makefile` that creates a brand new cluster to help with testing web-deployer. Feel free to use it too.

## Usage

```
make PROJECT=my-project ZONE=europe-west2-a CLUSTER=my-cluster
```

This will then log you into Google Cloud with `gcloud auth login` before setting up a new cluster and installing help within it with a secure service account.
