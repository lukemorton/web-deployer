package publish

import (
	"errors"
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/lukemorton/web-deployer/internal/log"
)

type VersionGateway interface {
	Exists(project string, name string, version string) (bool, error)
	Push(project string, name string, version string, dir string) error
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

func (g *versionGateway) Push(project string, name string, version string, dir string) error {
	err := g.ensureInstalled()

	if err != nil {
		return err
	}

	fullyQualifiedRepo := fullyQualifiedRepo(project, name, version)
	g.logger.Debugf("Building %s as %s...", dir, fullyQualifiedRepo)
	err = g.build(fullyQualifiedRepo, dir)

	if err != nil {
		return err
	}

	g.logger.Debugf("Pushing %s...", fullyQualifiedRepo)
	return runExecutable(g.logger.Writer(), "docker", "push", fullyQualifiedRepo)
}

func (g *versionGateway) ensureInstalled() (err error) {
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

func (g *versionGateway) build(fullyQualifiedRepo string, dir string) error {
	image, err := g.detect(dir)

	if err != nil {
		return err
	}

	return runExecutable(g.logger.Writer(), "s2i", "build", dir, image, fullyQualifiedRepo)
}

func (g *versionGateway) detect(dir string) (string, error) {
	files, err := ioutil.ReadDir(dir)

	if err != nil {
		return "", nil
	}

	for _, file := range files {
		g.logger.Debug(file.Name())

		if file.Name() == "config.ru" {
			return "centos/ruby-24-centos7", nil
		} else if strings.Contains(file.Name(), ".csproj") {
			return "registry.centos.org/dotnet/dotnet-20-centos7", nil
		}
	}

	return "", errors.New("Could not detect")
}

func repo(project string, name string) string {
	return fmt.Sprintf("gcr.io/%s/%s", project, name)
}

func fullyQualifiedRepo(project string, name string, version string) string {
	return fmt.Sprintf("gcr.io/%s/%s:%s", project, name, version)
}
