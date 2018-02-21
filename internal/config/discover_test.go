package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestDiscoveringWebDeployerYaml(t *testing.T) {
	config, err := Discover("../fixtures/ruby/")

	assert.NoError(t, err)
	assert.Equal(t, "doorman-1200", config.Kubernetes.Project)
}

func TestNonExistentDirectory(t *testing.T) {
	_, err := ReadFile("../fixtures/bob/")

	assert.Error(t, err)
}

func TestDirectoryWithoutWebDeployerYaml(t *testing.T) {
	_, err := ReadFile("../fixtures/")

	assert.Error(t, err)
}
