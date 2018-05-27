package publish

import (
	"errors"
	"testing"

	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
)

func TestPublishingSuccessfully(t *testing.T) {
	publisher, versionGateway := mockedPublisher()
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(false, nil)
	versionGateway.On("Push", "project", "app-staging", "v1", ".").Return(nil)

	assert.NoError(t, publisher.Publish("project", "app-staging", "v1", "."))

	versionGateway.AssertExpectations(t)
}

func TestPublishHandlesNotInstalled(t *testing.T) {
	publisher, versionGateway := mockedPublisher()
	s2iNotInstalledError := errors.New("s2i not installed")
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(false, s2iNotInstalledError)

	assert.Equal(t, s2iNotInstalledError, publisher.Publish("project", "app-staging", "v1", "."))

	versionGateway.AssertExpectations(t)
}

func TestPublishHandlesExistingTags(t *testing.T) {
	publisher, versionGateway := mockedPublisher()
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(true, nil)

	assert.Error(t, publisher.Publish("project", "app-staging", "v1", "."))

	versionGateway.AssertExpectations(t)
}

func TestPublishHandlesPushImageError(t *testing.T) {
	publisher, versionGateway := mockedPublisher()
	versionGateway.On("Exists", "project", "app-staging", "v1").Return(false, nil)
	pushError := errors.New("failed to push")
	versionGateway.On("Push", "project", "app-staging", "v1", ".").Return(pushError)

	assert.Equal(t, publisher.Publish("project", "app-staging", "v1", "."), pushError)

	versionGateway.AssertExpectations(t)
}

func mockedPublisher() (*publisher, *mockVersionGateway) {
  logger, _ := test.NewNullLogger()
	versionGateway := &mockVersionGateway{}
	return &publisher{logger: logger, versionGateway: versionGateway}, versionGateway
}
