package publish

import(
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

func (p *publisher) Publish(dir string) (err error) {
  err = p.validateExecutablesExist()
  if err != nil {
    return err
  }

  err = p.buildImage(dir)
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

func (p *publisher) buildImage(dir string) (err error) {
  appName, err := appNameFromDir(dir)
	if err != nil {
		return err
	}

  err = p.sourceToImageGateway.Build(dir, "centos/ruby-24-centos7", appName)
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
