package main

import (
	"errors"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/deploy"
)

var (
	deployUsage = `Deploy a version of your application.

In order to deploy your image to gcr.io run the following command. <dir> must
contain a web-deployer.yml file.

If the version has not already been published, it will be published before it is
deployed.

  web-deployer deploy <dir> <version>
`
	deployError = errors.New("Could not complete deploy.")
)

type deployRunner struct {
	dir        string
	app        string
	version    string
	k8sProject string
	logger     logrus.FieldLogger
}

func newDeployCmd(logger logrus.FieldLogger) *cobra.Command {
	runner := &deployRunner{logger: logger}

	cmd := &cobra.Command{
		Use:          "deploy <app> <version>",
		Short:        "Deploy a version of your application.",
		Long:         deployUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				runner.logger.Error("you must pass <app> and <version>")
				return deployError
			}

			dir, err := os.Getwd()
			if err != nil {
				return err
			}

			runner.dir = dir
			runner.app = args[0]
			runner.version = args[1]

			return runner.run()
		},
	}

	return cmd
}

func (runner *deployRunner) run() error {
	cfg, err := config.Discover(runner.dir)
	if err != nil {
		runner.logger.Errorf("Could not discover web-deployer.yml in %s", runner.dir)
		return deployError
	}

	appCfg, appIsDefined := cfg.Apps[runner.app]
	if appIsDefined == false {
		runner.logger.Errorf("Did not find `%s` app defined in web-deployer.yml", runner.app)
		return deployError
	}

	if len(cfg.Kubernetes.Project) == 0 {
		runner.logger.Error("Please specify a Kubernetes project in your web-deployer.yml")
		return deployError
	}

	runner.logger.Info("Deploying...")

	err = deploy.NewDeployer().Deploy(cfg.Kubernetes.Project, cfg.Kubernetes.Cluster, appCfg.Name, runner.version, runner.dir, appCfg.Hosts)
	if err != nil {
		runner.logger.Error(err)
		return deployError
	}

	runner.logger.Info("Deploy complete.")
	return nil
}
