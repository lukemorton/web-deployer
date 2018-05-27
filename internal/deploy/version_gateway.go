package deploy

type VersionGateway interface {
	EnsureInstalled() error
	Exists(project string, name string, version string) (bool, error)
	Deploy(project string, cluster string, name string, version string, hosts []string) error
}
