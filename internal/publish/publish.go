package publish

import (
	"fmt"
)

type publisher struct {
	sourceToImageGateway SourceToImageGateway
	gcloudGateway        GCloudGateway
}

func NewPublisher() *publisher {
	return &publisher{
		&sourceToImageGateway{},
		&gcloudGateway{},
	}
}

func (p *publisher) Publish(project string, name string, version string, dir string) (err error) {
	err = p.validateExecutablesExist()
	if err != nil {
		return err
	}

	repo := fmt.Sprintf("gcr.io/%s/%s", project, name)
	fullyQualifiedRepo := fmt.Sprintf("gcr.io/%s/%s:%s", project, name, version)

	err = p.validateImageDoesntExist(repo, version)
	if err != nil {
		return err
	}

	err = p.buildImage(dir, fullyQualifiedRepo)
	if err != nil {
		return err
	}

	err = p.pushImage(fullyQualifiedRepo)
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

func (p *publisher) validateImageDoesntExist(repo string, version string) error {
	exists, err := p.gcloudGateway.ImageTagExists(repo, version)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("Repository %s already has tag %s", repo, version)
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
