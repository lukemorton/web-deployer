package publish

import (
	"fmt"

	"github.com/lukemorton/web-deployer/internal/logger"
)

type Publisher interface {
	Publish(project string, name string, version string, dir string) (err error)
}

type publisher struct {
	logger         logger.Logger
	versionGateway VersionGateway
}

func NewPublisher(logger logger.Logger) *publisher {
	return &publisher{
		logger,
		&versionGateway{logger},
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
	p.logger.Info("Ensuring executables exist...")
	return p.versionGateway.EnsureInstalled()
}

func (p *publisher) validateVersionDoesntExist(project string, name string, version string) error {
	p.logger.Info("Ensuring version not already published...")
	exists, err := p.versionGateway.Exists(project, name, version)

	if err != nil {
		return err
	}

	if exists {
		p.logger.Errorf("Version already deployed project=%s name=%s version=%s", project, name, version)
		return fmt.Errorf("Version already deployed.")
	}

	return nil
}

func (p *publisher) push(project string, name string, version string, dir string) (err error) {
	p.logger.Info("Publishing version...")
	return p.versionGateway.Push(project, name, version, dir)
}
