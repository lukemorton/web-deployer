package main

import (
  "errors"
  "fmt"
  "io"

	"github.com/spf13/cobra"

  "github.com/lukemorton/web-deployer/internal/publish"
)


var publishUsage = `Publish an image of your application.

In order to push your image to gcr.io you should pass project:

  web-deployer publish --k8s-project <project> <dir>
`

type publishRunner struct {
  dir  string
  k8sProject string
	out  io.Writer
}

func newPublishCmd(out io.Writer) *cobra.Command {
  runner := &publishRunner{out: out}

	cmd := &cobra.Command{
		Use:          "publish <dir>",
		Short:        "Publish an image of your application.",
		Long:         publishUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
      if len(args) == 0 {
        return errors.New("you must pass <dir>")
      }

      runner.dir = args[0]
			return runner.run()
		},
	}

  cmd.Flags().StringVar(&runner.k8sProject, "k8s-project", "", "Kubernetes project to deploy to")
	return cmd
}

func (runner *publishRunner) run() error {
  if len(runner.k8sProject) == 0 {
    fmt.Fprintf(runner.out, "\n")
    return errors.New("Please specify a Kubernetes project with --k8s-project")
  }

  fmt.Fprintf(runner.out, "Publishing...")

  err := publish.NewPublisher().Publish(runner.k8sProject, runner.dir)
  if err != nil {
    fmt.Fprintf(runner.out, "\n")
    return err
  }

  fmt.Fprintf(runner.out, " done.")
  return nil
}
