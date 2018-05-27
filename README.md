# Web Deployer

For deploying a conventional web application to a conventional destination. Think of it as a convention over configuration approach to deployments.

## Supported deployments => platforms

We currently only support deploying to Google Kubernetes Engine platform. We will add additional target platforms in future such as AWS as well as more deployment applications.

- ✅ Ruby Rack application => Google Kubernetes Engine
- ✅ .NET application => Google Kubernetes Engine
- ❌ Node.js application => Google Kubernetes Engine
- ❌ Python application => Google Kubernetes Engine
- ❌ PHP application => Google Kubernetes Engine

## Installation

You can use our docker image that already has prerequisites and the latest `web-deployer` installed. All you need to do is make sure you have [Docker](https://docs.docker.com/install/) installed.

```
docker run --rm -ti lukemorton/web-deployer-tools ash
```

Or read our [Advanced Install](#advanced-install) guide.

## Usage

**Configuring your deployment**

We need to configure where on GKE we are going to deploy and also what deployments we want to create.

Create a `web-deployer.yml` next to your `config.ru` file with the following contents:

``` yml
gcloud:
  project: my-project
  zone: europe-west2-a
  cluster: my-cluster

deployments:
  staging:
    name: ruby-sample-app-staging
    hosts:
      - ruby-sample-app-staging.local
  production:
    name: ruby-sample-app-production
    hosts:
      - ruby-sample-app-production.local
```

**Bootstrapping your deployments on GKE**

The following command will then loop over each deployment and create an example application.

```
web-deployer bootstrap
```

If you then add a record into your `/etc/hosts` file mapping the hosts to the IP returned by the bootstrap command.

**Deploying**

You can then deploy:

```
web-deployer deploy staging v1.1
```

The staging deployment will then be updated to serve your `config.ru`.

## Advanced Install

**Prerequisites**

 - Go (install from https://golang.org/)
 - Docker (install from https://www.docker.com/)
 - s2i (`brew install source-to-image`)
 - gcloud (`brew install caskroom/cask/google-cloud-sdk`)
 - kubectl (`brew install kubectl`)
 - helm (`brew install kubernetes-helm`)

**Compiling web-deployer**

You will need a Go environment installed. Then run:

```
go install github.com/lukemorton/web-deployer/cmd/web-deployer
```
