package publish

import(
	"github.com/stretchr/testify/mock"
)

type mockSourceToImageGateway struct {
	mock.Mock
}

func (g *mockSourceToImageGateway) EnsureInstalled() error {
	args := g.Called()
	return args.Error(0)
}

func (g *mockSourceToImageGateway) Build(dir string, image string, tag string) error {
  args := g.Called(dir, image, tag)
	return args.Error(0)
}
