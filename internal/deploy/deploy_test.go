package deploy

import (
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestDeployingAndPublishingVersionSuccessfully(t *testing.T) {
	deployer, publisher, versionGateway := mockedDeployer()
	versionGateway.On("EnsureInstalled").Return(nil)
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(false, nil)
	publisher.On("Publish", "project", "app-staging", "v1", ".").Return(nil)
	versionGateway.On("Deploy", "project", "cluster", "app-staging", "v1", []string{"cool.com"}).Return(nil)

	assert.NoError(t, deployer.Deploy("project", "cluster", "app-staging", "v1", ".", []string{"cool.com"}))

	versionGateway.AssertExpectations(t)
	publisher.AssertExpectations(t)
}

func TestDeployingPublishedVersionSuccessfully(t *testing.T) {
	deployer, publisher, versionGateway := mockedDeployer()
	versionGateway.On("EnsureInstalled").Return(nil)
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(true, nil)
	versionGateway.On("Deploy", "project", "cluster", "app-staging", "v1", []string{"cool.com"}).Return(nil)

	assert.NoError(t, deployer.Deploy("project", "cluster", "app-staging", "v1", ".", []string{"cool.com"}))

	versionGateway.AssertExpectations(t)
	publisher.AssertExpectations(t)
}

func mockedDeployer() (*deployer, *mockPublisher, *mockVersionGateway) {
	logger, _ := test.NewNullLogger()
	publisher := &mockPublisher{}
	versionGateway := &mockVersionGateway{}
	return &deployer{logger, publisher, versionGateway}, publisher, versionGateway
}
