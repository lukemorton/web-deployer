package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/publish"
)

var publishUsage = `Publish an image of your application.

In order to push your image to gcr.io run the following command. <dir> must
contain a web-deployer.yml file.

  web-deployer publish <dir> <version>
`

type publishRunner struct {
	dir        string
	app        string
	version    string
	k8sProject string
	out        io.Writer
}

func newPublishCmd(out io.Writer) *cobra.Command {
	runner := &publishRunner{out: out}

	cmd := &cobra.Command{
		Use:          "publish <app> <version>",
		Short:        "Publish an image of your application.",
		Long:         publishUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				return errors.New("you must pass <app> and <version>")
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			runner.dir = dir
			runner.app = args[0]
			runner.version = args[1]

			err = runner.run()
			if err != nil {
				fmt.Fprintf(runner.out, "\n")
			}
			return err
		},
	}

	return cmd
}

func (runner *publishRunner) run() error  {
	cfg, err := config.Discover(runner.dir)
	if err != nil {
		return fmt.Errorf("Could not discover web-deployer.yml in %s", runner.dir)
	}

	appCfg, appIsDefined := cfg.Apps[runner.app]
	if appIsDefined == false {
		return fmt.Errorf("Did not find `%s` app defined in web-deployer.yml", runner.app)
	}

	if len(cfg.Kubernetes.Project) == 0 {
		return errors.New("Please specify a Kubernetes project in your web-deployer.yml")
	}

	fmt.Fprintf(runner.out, "Publishing...")

	err = publish.NewPublisher().Publish(cfg.Kubernetes.Project, appCfg.Name, runner.version, runner.dir)
	if err != nil {
		return err
	}

	fmt.Fprintf(runner.out, " done.")
	return nil
}
