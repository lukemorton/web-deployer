package publish

import (
	"fmt"
	"strings"

	"github.com/lukemorton/web-deployer/internal/logger"
)

type VersionGateway interface {
	EnsureInstalled() error
	Exists(project string, name string, version string) (bool, error)
	Push(project string, name string, version string, dir string) error
}

type versionGateway struct {
	logger logger.Logger
}

func (g *versionGateway) EnsureInstalled() (err error) {
	g.logger.Debug("Looking for docker executable...")
	err = ensureExecutableInstalled("docker")
	if err != nil {
		return err
	}

	g.logger.Debug("Looking for gcloud executable...")
	err = ensureExecutableInstalled("gcloud")
	if err != nil {
		return err
	}

	g.logger.Debug("Looking for s2i executable...")
	err = ensureExecutableInstalled("s2i")
	if err != nil {
		return err
	}

	return nil
}

func (g *versionGateway) Exists(project string, name string, version string) (bool, error) {
	repo := repo(project, name)
	g.logger.Debugf("Looking for %s version in %s...", version, repo)
	out, err := runExecutableAndReturnOutput("gcloud", "container", "images", "list-tags", repo, "--filter", version, "--format", "json")
	return strings.TrimSpace(string(out)) != "[]", err
}

func (g *versionGateway) Push(project string, name string, version string, dir string) error {
	fullyQualifiedRepo := fullyQualifiedRepo(project, name, version)
	g.logger.Debugf("Building %s as %s...", dir, repo)
	err := build(fullyQualifiedRepo, dir)

	if err != nil {
		return err
	}

	g.logger.Info("Pushing %s...", repo)
	return runExecutable("gcloud", "docker", "--", "push", fullyQualifiedRepo)
}

func repo(project string, name string) string {
	return fmt.Sprintf("gcr.io/%s/%s", project, name)
}

func fullyQualifiedRepo(project string, name string, version string) string {
	return fmt.Sprintf("gcr.io/%s/%s:%s", project, name, version)
}

func build(fullyQualifiedRepo string, dir string) error {
	return runExecutable("s2i", "build", dir, "centos/ruby-24-centos7", fullyQualifiedRepo)
}
