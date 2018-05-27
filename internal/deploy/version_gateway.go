package deploy

import (
	"fmt"
	"strings"

	"github.com/lukemorton/web-deployer/internal/log"
)

type VersionGateway interface {
	Exists(project string, name string, version string) (bool, error)
	Deploy(project string, cluster string, name string, version string, hosts []string) error
}

type versionGateway struct {
	logger log.Logger
}

func (g *versionGateway) Exists(project string, name string, version string) (bool, error) {
	err := g.ensureInstalled()

	if err != nil {
		return false, err
	}

	repo := repo(project, name)
	g.logger.Debugf("Looking for %s version in %s...", version, repo)
	out, err := runExecutableAndReturnOutput(g.logger.Writer(), "gcloud", "container", "images", "list-tags", repo, "--filter", version, "--format", "json")
	g.logger.Debugf("Versions found: %s", out)
	return strings.TrimSpace(string(out)) != "[]", err
}

func (g *versionGateway) Deploy(project string, cluster string, name string, version string, hosts []string) (err error) {
	err = g.ensureInstalled()

	if err != nil {
		return err
	}

	err = g.loadClusterCredentials(cluster)
	if err != nil {
		return err
	}

	err = g.helmInit()
	if err != nil {
		return err
	}

	return g.helmUpgrade(project, name, version, hosts)
}

func (g *versionGateway) ensureInstalled() (err error) {
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

func (g *versionGateway) loadClusterCredentials(cluster string) error {
	return runExecutable(g.logger.Writer(), "gcloud", "container", "clusters", "get-credentials", cluster)
}

func (g *versionGateway) helmInit() error {
	return runExecutable(g.logger.Writer(), "helm", "init", "--client-only")
}

func (g *versionGateway) helmUpgrade(project string, name string, version string, hosts []string) error {
	repo := repo(project, name)
	return runExecutable(
		g.logger.Writer(),
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
