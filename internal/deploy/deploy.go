package deploy

import (
	"github.com/lukemorton/web-deployer/internal/logger"
	"github.com/lukemorton/web-deployer/internal/publish"
)

type deployer struct {
	publisher      publish.Publisher
	versionGateway VersionGateway
}

func NewDeployer(logger logger.Logger) *deployer {
	return &deployer{
		publisher: publish.NewPublisher(logger),
	}
}

func (d *deployer) Deploy(project string, cluster string, name string, version string, dir string, hosts []string) (err error) {
	err = d.validateExecutablesExist()
	if err != nil {
		return err
	}

	versionExists, err := d.versionExists(project, name, version)

	if err != nil {
		return err
	}

	if versionExists == false {
		err = d.publish(project, name, version, dir)

		if err != nil {
			return err
		}
	}

	return d.deploy(project, cluster, name, version, hosts)
}

func (d *deployer) validateExecutablesExist() error {
	return d.versionGateway.EnsureInstalled()
}

func (d *deployer) versionExists(project string, name string, version string) (bool, error) {
	return d.versionGateway.Exists(project, name, version)
}

func (d *deployer) publish(project string, name string, version string, dir string) error {
	return d.publisher.Publish(project, name, version, dir)
}

func (d *deployer) deploy(project string, cluster string, name string, version string, hosts []string) error {
	return d.versionGateway.Deploy(project, cluster, name, version, hosts)
}
