package publish

type GCloudGateway interface {
  EnsureInstalled() error
  LoadClusterCredentials(cluster string) error
  PushImage(repo string) error
}

type gcloudGateway struct {
}

func (g *gcloudGateway) EnsureInstalled() (err error) {
  err = ensureExecutableInstalled("docker")
  if err != nil {
    return err
  }

  err = ensureExecutableInstalled("gcloud")
  if err != nil {
    return err
  }

  return nil
}

func (g *gcloudGateway) LoadClusterCredentials(cluster string) error {
  return runExecutable("gcloud", "clusters", "get-credentials", cluster)
}

func (g *gcloudGateway) PushImage(repo string) error {
  return runExecutable("gcloud", "docker", "--", "push", repo)
}
