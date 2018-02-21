package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReadingFile(t *testing.T) {
	config, err := ReadFile("../fixtures/ruby/web-deployer.yml")

	assert.NoError(t, err)
	assert.Equal(t, "doorman-1200", config.Kubernetes.Project)
}

func TestNonExistentReadingFile(t *testing.T) {
	_, err := ReadFile("../fixtures/ruby/bob.yml")

	assert.Error(t, err)
}
