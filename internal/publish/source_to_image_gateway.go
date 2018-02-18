package publish

type SourceToImageGateway interface {
  Installed() bool
  Build(dir string, image string, tag string) error
}

type sourceToImageGateway struct {

}

func (g *sourceToImageGateway) Installed() bool {
  return isExecutableInstalled("s2i")
}

func (g *sourceToImageGateway) Build(dir string, image string, tag string) error {
  return runExecutable("s2i", "build", dir, image, tag)
}
