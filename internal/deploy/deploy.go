package deploy

type deployer struct {
}

func NewDeployer() *deployer {
	return &deployer{}
}

func (p *deployer) Deploy(project string, name string, version string, dir string) (err error) {
	return nil
}
