package deploy

import (
	"github.com/lukemorton/web-deployer/internal/log"
	"github.com/lukemorton/web-deployer/internal/publish"
)

type deployer struct {
	logger         log.Logger
	publisher      publish.Publisher
	versionGateway VersionGateway
}

func NewDeployer(logger log.Logger) *deployer {
	return &deployer{
		logger: logger,
		publisher: publish.NewPublisher(logger),
		versionGateway: &versionGateway{logger},
	}
}

func (d *deployer) Deploy(project string, zone string, cluster string, name string, version string, dir string, hosts []string) (err error) {
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

	return d.deploy(project, zone, cluster, name, version, hosts)
}

func (d *deployer) versionExists(project string, name string, version string) (exists bool, err error) {
	d.logger.Info("Deciding whether to publish...")
	exists, err = d.versionGateway.Exists(project, name, version)

	if exists {
		d.logger.Info("Already published...")
	}

	return
}

func (d *deployer) publish(project string, name string, version string, dir string) error {
	return d.publisher.Publish(project, name, version, dir)
}

func (d *deployer) deploy(project string, zone string, cluster string, name string, version string, hosts []string) error {
	d.logger.Info("Deploying version...")
	return d.versionGateway.Deploy(project, zone, cluster, name, version, hosts)
}
