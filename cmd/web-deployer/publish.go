package main

import (
	"errors"
	"fmt"
	"io"

	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/publish"
)

var publishUsage = `Publish an image of your application.

In order to push your image to gcr.io run the following command. <dir> must
contain a web-deployer.yml file.

  web-deployer publish <dir>
`

type publishRunner struct {
	dir        string
	k8sProject string
	out        io.Writer
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

	return cmd
}

func (runner *publishRunner) run() error  {
	cfg, err := config.Discover(runner.dir)
	if err != nil {
		fmt.Fprintf(runner.out, "\n")
		return errors.New("Could not discover web-deployer.yml")
	}

	if len(cfg.Kubernetes.Project) == 0 {
		fmt.Fprintf(runner.out, "\n")
		return errors.New("Please specify a Kubernetes project in your web-deployer.yml")
	}

	fmt.Fprintf(runner.out, "Publishing...")

	err = publish.NewPublisher().Publish(cfg.Kubernetes.Project, runner.dir)
	if err != nil {
		fmt.Fprintf(runner.out, "\n")
		return err
	}

	fmt.Fprintf(runner.out, " done.")
	return nil
}
