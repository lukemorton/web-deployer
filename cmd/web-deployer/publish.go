package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/log"
	"github.com/lukemorton/web-deployer/internal/publish"
)

var (
	publishUsage = `Publish a version of your application.

In order to push your image to gcr.io run the following command. <deployment> must
contain a web-deployer.yml file.

  web-deployer publish <deployment> <version>
`

	publishError = errors.New("Could not complete publishing.")
)

type publishRunner struct {
	dir        string
	deployment string
	version    string
	logger     log.Logger
}

func newPublishCmd(logger log.Logger) *cobra.Command {
	runner := &publishRunner{logger: logger}

	cmd := &cobra.Command{
		Use:          "publish <deployment> <version>",
		Short:        "Publish a version of your application.",
		Long:         publishUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				logger.Error("you must pass <deployment> and <version>")
				return publishError
			}

			dir, err := os.Getwd()
			if err != nil {
				logger.Error(err)
				return publishError
			}

			runner.dir = dir
			runner.deployment = args[0]
			runner.version = args[1]

			return runner.run()
		},
	}

	return cmd
}

func (runner *publishRunner) run() error {
	cfg, err := config.Discover(runner.dir)
	if err != nil {
		runner.logger.Errorf("Could not discover web-deployer.yml in %s", runner.dir)
		return publishError
	}

	deployment, deploymentIsDefined := cfg.Deployments[runner.deployment]
	if deploymentIsDefined == false {
		runner.logger.Errorf("Did not find `%s` deployment defined in web-deployer.yml", runner.deployment)
		return publishError
	}

	if len(cfg.GCloud.Project) == 0 {
		runner.logger.Errorf("Please specify a GCloud project in your web-deployer.yml")
		return publishError
	}

	runner.logger.Info("Publishing...")

	err = publish.NewPublisher(runner.logger).Publish(cfg.GCloud.Project, deployment.Name, runner.version, runner.dir)
	if err != nil {
		return publishError
	}

	runner.logger.Info("Publishing complete.")
	return nil
}
