package deploy

import (
	"github.com/stretchr/testify/mock"
)

type mockVersionGateway struct {
	mock.Mock
}

func (g *mockVersionGateway) Exists(project string, name string, version string) (bool, error) {
	args := g.Called(project, name, version)
	return args.Bool(0), args.Error(1)
}

func (g *mockVersionGateway) Deploy(project string, cluster string, name string, version string, hosts []string) error {
	args := g.Called(project, cluster, name, version, hosts)
	return args.Error(0)
}
