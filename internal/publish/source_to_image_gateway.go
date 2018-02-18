package publish

type SourceToImageGateway interface {
	EnsureInstalled() error
	Build(dir string, image string, tag string) error
}

type sourceToImageGateway struct {
}

func (g *sourceToImageGateway) EnsureInstalled() error {
	return ensureExecutableInstalled("s2i")
}

func (g *sourceToImageGateway) Build(dir string, image string, tag string) error {
	return runExecutable("s2i", "build", dir, image, tag)
}
