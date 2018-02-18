package publish

import(
  "fmt"
  "path/filepath"
)

type publisher struct {
  sourceToImageGateway SourceToImageGateway
  gcloudGateway GCloudGateway
}

func NewPublisher() *publisher {
  return &publisher{
    &sourceToImageGateway{},
    &gcloudGateway{},
  }
}

func (p *publisher) Publish(project string, dir string) (err error) {
  err = p.validateExecutablesExist()
  if err != nil {
    return err
  }

  appName, err := appNameFromDir(dir)
	if err != nil {
		return err
	}

  repo := fmt.Sprintf("gcr.io/%s/%s", project, appName)

  err = p.buildImage(dir, repo)
  if err != nil {
    return err
  }

  err = p.pushImage(repo)
  if err != nil {
    return err
  }

  return nil
}

func (p *publisher) validateExecutablesExist() (err error) {
  err = p.sourceToImageGateway.EnsureInstalled()
  if err != nil {
    return err
  }

  err = p.gcloudGateway.EnsureInstalled()
  if err != nil {
    return err
  }

  return nil
}

func (p *publisher) buildImage(dir string, repo string) (err error) {
  err = p.sourceToImageGateway.Build(dir, "centos/ruby-24-centos7", repo)
	if err != nil {
		return err
	}

  return nil
}

func (p *publisher) pushImage(repo string) (err error) {
  err = p.gcloudGateway.PushImage(repo)
	if err != nil {
		return err
	}

  return nil
}

func appNameFromDir(dir string) (string, error) {
  path, err := filepath.Abs(dir)

  if err != nil {
    return path, err
  }

  return filepath.Base(path), nil
}
