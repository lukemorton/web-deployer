package config

type Config struct {
  Kubernetes KubernetesConfig `yaml:"k8s"`
  Apps       map[string]AppConfig
}

type KubernetesConfig struct {
  Project string
  Zone    string
  Cluster string
}

type AppConfig struct {
  Name   string
  Hosts  []string
  Charts []map[string]map[string]string
}
