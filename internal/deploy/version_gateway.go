package deploy

import (
	"fmt"
	"strings"

	"github.com/lukemorton/web-deployer/internal/log"
)

type VersionGateway interface {
	EnsureInstalled() error
	Exists(project string, name string, version string) (bool, error)
	Deploy(project string, cluster string, name string, version string, hosts []string) error
}

type versionGateway struct {
	logger log.Logger
}

func (g *versionGateway) EnsureInstalled() (err error) {
	g.logger.Debug("Looking for gcloud executable...")
	err = ensureExecutableInstalled("gcloud")
	if err != nil {
		return err
	}

	g.logger.Debug("Looking for helm executable...")
	err = ensureExecutableInstalled("helm")
	if err != nil {
		return err
	}

	return nil
}

func (g *versionGateway) Exists(project string, name string, version string) (bool, error) {
	repo := repo(project, name)
	out, err := runExecutableAndReturnOutput("gcloud", "container", "images", "list-tags", repo, "--filter", version, "--format", "json")
	return strings.TrimSpace(string(out)) != "[]", err
}

func (g *versionGateway) Deploy(project string, cluster string, name string, version string, hosts []string) (err error) {
	err = loadClusterCredentials(cluster)
	if err != nil {
		return err
	}

	err = helmInit()
	if err != nil {
		return err
	}

	return helmUpgrade(project, name, version, hosts)
}

func loadClusterCredentials(cluster string) error {
	return runExecutable("gcloud", "container", "clusters", "get-credentials", cluster)
}

func helmInit() error {
	return runExecutable("helm", "init", "--client-only")
}

func helmUpgrade(project string, name string, version string, hosts []string) error {
	repo := repo(project, name)
	return runExecutable(
		"helm",
		"upgrade",
		name,
		"http://web-deployer-charts.storage.googleapis.com/web-app-0.1.0.tgz",
		"--install",
		"--set", fmt.Sprintf("image.repository=%s", repo),
		"--set", fmt.Sprintf("image.tag=%s", version),
		"--set", fmt.Sprintf("ingress.hosts={%s}", strings.Join(hosts, ",")),
	)
}

func repo(project string, name string) string {
	return fmt.Sprintf("gcr.io/%s/%s", project, name)
}
