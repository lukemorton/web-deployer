package publish

import(
  "fmt"
  "path/filepath"
)

type publisher struct {
  sourceToImageGateway SourceToImageGateway
}

func NewPublisher() *publisher {
  return &publisher{
    &sourceToImageGateway{},
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
  executables := []string{"docker", "gcloud"}

  for _, executable := range executables {
    if isExecutableInstalled(executable) != true {
      return fmt.Errorf("Could not find %s executable, is it installed?", executable)
    }
  }

  err = validateExecutableExists(p.sourceToImageGateway)
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

type ExecutableInstalled interface {
  Installed() bool
}

func validateExecutableExists(executable ExecutableInstalled) error {
  if executable.Installed() != true {
    return fmt.Errorf("Could not find %s executable, is it installed?", executable)
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
