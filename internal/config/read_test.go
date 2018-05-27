package config

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParsingConfig(t *testing.T) {
	yaml := `
    gcloud: { project: "cool" }
    deployments:
      staging:
        name: cool-staging
        charts:
          - gcloud-sqlproxy:
              cloudsql.instance: staging
        hosts:
          - staging.cool.com
  `
	config, err := Read([]byte(yaml))

	assert.NoError(t, err)
	assert.Equal(t, "cool", config.GCloud.Project)
	assert.Equal(t, "cool-staging", config.Deployments["staging"].Name)
	assert.Equal(t, "staging.cool.com", config.Deployments["staging"].Hosts[0])
	assert.Equal(t, "staging", config.Deployments["staging"].Charts[0]["gcloud-sqlproxy"]["cloudsql.instance"])
}

func TestParsingKeyThatDoesNotExist(t *testing.T) {
	yaml := `bob: { project: "cool" }`
	_, err := Read([]byte(yaml))

	assert.NoError(t, err)
}
