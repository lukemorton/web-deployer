package publish

import (
	"fmt"
)

type Publisher interface {
	Publish(project string, name string, version string, dir string) (err error)
}

type publisher struct {
	versionGateway VersionGateway
}

func NewPublisher() *publisher {
	return &publisher{
		&versionGateway{},
	}
}

func (p *publisher) Publish(project string, name string, version string, dir string) (err error) {
	err = p.validateExecutablesExist()
	if err != nil {
		return err
	}

	err = p.validateVersionDoesntExist(project, name, version)
	if err != nil {
		return err
	}

	err = p.push(project, name, version, dir)
	if err != nil {
		return err
	}

	return nil
}

func (p *publisher) validateExecutablesExist() error {
	return p.versionGateway.EnsureInstalled()
}

func (p *publisher) validateVersionDoesntExist(project string, name string, version string) error {
	exists, err := p.versionGateway.Exists(project, name, version)

	if err != nil {
		return err
	}

	if exists {
		return fmt.Errorf("Version already deployed project=%s name=%s version=%s", project, name, version)
	}

	return nil
}

func (p *publisher) push(project string, name string, version string, dir string) (err error) {
	return p.versionGateway.Push(project, name, version, dir)
}
