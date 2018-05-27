package deploy

import (
	"github.com/stretchr/testify/mock"
)

type mockPublisher struct {
	mock.Mock
}

func (p *mockPublisher) Publish(project string, name string, version string, dir string) (err error) {
	args := p.Called(project, name, version, dir)
	return args.Error(0)
}
