package publish

import (
  "fmt"
	"strings"
)

type VersionGateway interface {
	EnsureInstalled() error
	Exists(project string, name string, version string) (bool, error)
	Push(project string, name string, version string, dir string) error
}

type versionGateway struct {
}

func (g *versionGateway) EnsureInstalled() (err error) {
	err = ensureExecutableInstalled("docker")
	if err != nil {
		return err
	}

	err = ensureExecutableInstalled("gcloud")
	if err != nil {
		return err
	}

	err = ensureExecutableInstalled("s2i")
	if err != nil {
		return err
	}

	return nil
}

func (g *versionGateway) LoadClusterCredentials(cluster string) error {
	return runExecutable("gcloud", "clusters", "get-credentials", cluster)
}

func (g *versionGateway) Exists(project string, name string, version string) (bool, error) {
  repo := repo(project, name)
	out, err := runExecutableAndReturnOutput("gcloud", "container", "images", "list-tags", repo, "--filter", version, "--format", "json")
	return strings.TrimSpace(string(out)) != "[]", err
}

func (g *versionGateway) Push(project string, name string, version string, dir string) error {
  fullyQualifiedRepo := fullyQualifiedRepo(project, name, version)
  err := build(fullyQualifiedRepo, dir)

  if err != nil {
    return err
  }

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
