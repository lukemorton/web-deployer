package publish

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPublishingSuccessfully(t *testing.T) {
	publisher, sourceToImageGateway, gcloudGateway := mockedPublisher()
	sourceToImageGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("ImageTagExists", "gcr.io/project/app-staging", "v1").Return(false, nil)
	sourceToImageGateway.On("Build", ".", "centos/ruby-24-centos7", "gcr.io/project/app-staging:v1").Return(nil)
	gcloudGateway.On("PushImage", "gcr.io/project/app-staging:v1").Return(nil)

	assert.NoError(t, publisher.Publish("project", "app-staging", "v1", "."))

	sourceToImageGateway.AssertExpectations(t)
	gcloudGateway.AssertExpectations(t)
}

func TestPublishHandlesMissingSourceToImageExecutable(t *testing.T) {
	publisher, sourceToImageGateway, _ := mockedPublisher()
	s2iNotInstalledError := errors.New("s2i not installed")
	sourceToImageGateway.On("EnsureInstalled").Return(s2iNotInstalledError)

	assert.Equal(t, s2iNotInstalledError, publisher.Publish("project", "app-staging", "v1", "."))

	sourceToImageGateway.AssertExpectations(t)
}

func TestPublishHandlesMissingGCloudExecutable(t *testing.T) {
	publisher, sourceToImageGateway, gcloudGateway := mockedPublisher()
	gcloudNotInstalledError := errors.New("gcloud not installed")
	sourceToImageGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("EnsureInstalled").Return(gcloudNotInstalledError)

	assert.Equal(t, gcloudNotInstalledError, publisher.Publish("project", "app-staging", "v1", "."))

	gcloudGateway.AssertExpectations(t)
}

func TestPublishHandlesExistingTags(t *testing.T) {
	publisher, sourceToImageGateway, gcloudGateway := mockedPublisher()
	sourceToImageGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("ImageTagExists", "gcr.io/project/app-staging", "v1").Return(true, nil)

	assert.Error(t, publisher.Publish("project", "app-staging", "v1", "."))

	sourceToImageGateway.AssertExpectations(t)
	gcloudGateway.AssertExpectations(t)
}

func TestPublishHandlesBuildError(t *testing.T) {
	publisher, sourceToImageGateway, gcloudGateway := mockedPublisher()
	sourceToImageGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("ImageTagExists", "gcr.io/project/app-staging", "v1").Return(false, nil)
	buildError := errors.New("failed to build")
	sourceToImageGateway.On("Build", ".", "centos/ruby-24-centos7", "gcr.io/project/app-staging:v1").Return(buildError)

	assert.Equal(t, publisher.Publish("project", "app-staging", "v1", "."), buildError)

	sourceToImageGateway.AssertExpectations(t)
	gcloudGateway.AssertExpectations(t)
}

func TestPublishHandlesPushImageError(t *testing.T) {
	publisher, sourceToImageGateway, gcloudGateway := mockedPublisher()
	sourceToImageGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("EnsureInstalled").Return(nil)
	gcloudGateway.On("ImageTagExists", "gcr.io/project/app-staging", "v1").Return(false, nil)
	sourceToImageGateway.On("Build", ".", "centos/ruby-24-centos7", "gcr.io/project/app-staging:v1").Return(nil)
	pushImageError := errors.New("failed to push")
	gcloudGateway.On("PushImage", "gcr.io/project/app-staging:v1").Return(pushImageError)

	assert.Equal(t, publisher.Publish("project", "app-staging", "v1", "."), pushImageError)

	sourceToImageGateway.AssertExpectations(t)
	gcloudGateway.AssertExpectations(t)
}

func mockedPublisher() (*publisher, *mockSourceToImageGateway, *mockGCloudGateway) {
	sourceToImageGateway := &mockSourceToImageGateway{}
	gcloudGateway := &mockGCloudGateway{}
	return &publisher{sourceToImageGateway, gcloudGateway}, sourceToImageGateway, gcloudGateway
}
