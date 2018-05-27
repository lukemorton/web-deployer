FROM alpine as s2i

RUN apk add --update wget && \
  wget https://github.com/openshift/source-to-image/releases/download/v1.1.8/source-to-image-v1.1.8-e3140d01-linux-386.tar.gz && \
  tar -xf source-to-image-v1.1.8-e3140d01-linux-386.tar.gz && \
  mv s2i /usr/local/bin

FROM google/cloud-sdk:alpine

RUN apk add --update make docker && \
    gcloud components install kubectl

COPY --from=s2i /usr/local/bin/s2i /usr/local/bin/
COPY --from=lachlanevenson/k8s-helm /usr/local/bin/helm /usr/local/bin/
ADD dist/web-deployer /usr/local/bin/
