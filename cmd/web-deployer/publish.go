package main

import (
	"errors"
	"os"

	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/config"
	"github.com/lukemorton/web-deployer/internal/logger"
	"github.com/lukemorton/web-deployer/internal/publish"
)

var (
	publishUsage = `Publish a version of your application.

In order to push your image to gcr.io run the following command. <dir> must
contain a web-deployer.yml file.

  web-deployer publish <dir> <version>
`

	publishError = errors.New("Could not complete publishing.")
)

type publishRunner struct {
	dir        string
	app        string
	version    string
	k8sProject string
	logger     logger.Logger
}

func newPublishCmd(logger logger.Logger) *cobra.Command {
	runner := &publishRunner{logger: logger}

	cmd := &cobra.Command{
		Use:          "publish <app> <version>",
		Short:        "Publish a version of your application.",
		Long:         publishUsage,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 2 {
				logger.Error("you must pass <app> and <version>")
				return publishError
			}

			dir, err := os.Getwd()
			if err != nil {
				logger.Error(err)
				return publishError
			}

			runner.dir = dir
			runner.app = args[0]
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

	appCfg, appIsDefined := cfg.Apps[runner.app]
	if appIsDefined == false {
		runner.logger.Infof("Did not find `%s` app defined in web-deployer.yml", runner.app)
		return publishError
	}

	if len(cfg.Kubernetes.Project) == 0 {
		runner.logger.Infof("Please specify a Kubernetes project in your web-deployer.yml")
		return publishError
	}

	runner.logger.Info("Publishing...")

	err = publish.NewPublisher(runner.logger).Publish(cfg.Kubernetes.Project, appCfg.Name, runner.version, runner.dir)
	if err != nil {
		return publishError
	}

	runner.logger.Info("Publishing complete.")
	return nil
}
