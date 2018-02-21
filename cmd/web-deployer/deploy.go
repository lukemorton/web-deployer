package main

import (
	"errors"
	"fmt"
	"io"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/deploy"
)

var deployUsage = `Deploy a version of your application.

In order to deploy your image to gcr.io run the following command. <dir> must
contain a web-deployer.yml file.

If the version has not already been published, it will be published before it is
deployed.

  web-deployer deploy <dir> <version>
`

type deployRunner struct {
	dir        string
	app        string
	version    string
	k8sProject string
	out        io.Writer
}

func newDeployCmd(out io.Writer) *cobra.Command {
	runner := &deployRunner{out: out}

	cmd := &cobra.Command{
		Use:          "deploy <app> <version>",
		Short:        "Deploy a version of your application.",
		Long:         deployUsage,
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

func (runner *deployRunner) run() error {
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

	fmt.Fprintf(runner.out, "Deploying...")

	err = deploy.NewDeployer().Deploy(cfg.Kubernetes.Project, appCfg.Name, runner.version, runner.dir)
	if err != nil {
		return err
	}

	fmt.Fprintf(runner.out, " done.")
	return nil
}
