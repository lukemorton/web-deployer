package publish

import(
	"strings"
)

type GCloudGateway interface {
	EnsureInstalled() error
	LoadClusterCredentials(cluster string) error
	ImageTagExists(repo string, tag string) (bool, error)
	PushImage(repo string) error
}

type gcloudGateway struct {
}

func (g *gcloudGateway) EnsureInstalled() (err error) {
	err = ensureExecutableInstalled("docker")
	if err != nil {
		return err
	}

	err = ensureExecutableInstalled("gcloud")
	if err != nil {
		return err
	}

	return nil
}

func (g *gcloudGateway) LoadClusterCredentials(cluster string) error {
	return runExecutable("gcloud", "clusters", "get-credentials", cluster)
}

func (g *gcloudGateway) ImageTagExists(repo string, tag string) (bool, error) {
	out, err := runExecutableAndReturnOutput("gcloud", "container", "images", "list-tags", repo, "--filter", tag, "--format", "json")
	return strings.TrimSpace(string(out)) != "[]", err
}

func (g *gcloudGateway) PushImage(repo string) error {
	return runExecutable("gcloud", "docker", "--", "push", repo)
}
