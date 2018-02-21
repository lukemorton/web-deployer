package publish

import (
	"github.com/stretchr/testify/mock"
)

type mockGCloudGateway struct {
	mock.Mock
}

func (g *mockGCloudGateway) EnsureInstalled() error {
	args := g.Called()
	return args.Error(0)
}

func (g *mockGCloudGateway) LoadClusterCredentials(cluster string) error {
	return nil
}

func (g *mockGCloudGateway) ImageTagExists(repo string, tag string) (bool, error) {
	args := g.Called(repo, tag)
	return args.Bool(0), args.Error(1)
}

func (g *mockGCloudGateway) PushImage(repo string) error {
	args := g.Called(repo)
	return args.Error(0)
}
