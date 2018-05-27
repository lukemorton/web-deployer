package config

type Config struct {
	GCloud      GCloudConfig
	Deployments map[string]DeploymentConfig
}

type GCloudConfig struct {
	Project string
	Zone    string
	Cluster string
}

type DeploymentConfig struct {
	Name   string
	Hosts  []string
	Charts []map[string]map[string]string
}
