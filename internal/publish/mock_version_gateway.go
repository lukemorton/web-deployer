package publish

import (
	"github.com/stretchr/testify/mock"
)

type mockVersionGateway struct {
	mock.Mock
}

func (g *mockVersionGateway) EnsureInstalled() error {
	args := g.Called()
	return args.Error(0)
}

func (g *mockVersionGateway) Exists(project string, name string, version string) (bool, error) {
	args := g.Called(project, name, version)
	return args.Bool(0), args.Error(1)
}

func (g *mockVersionGateway) Push(project string, name string, version string, dir string) error {
	args := g.Called(project, name, version, dir)
	return args.Error(0)
}
