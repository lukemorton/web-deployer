package main

import (
	"io"
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/lukemorton/web-deployer/internal/log"
)

var globalUsage = `Web Deployer for Kubernetes.

Before getting started ensure you have activated a service account for
Kubernetes and have gotten credentials for a cluster.

	gcloud auth activate-service-account --key-file gcp-key.json
	gcloud container clusters get-credentials --project $(PROJECT) --zone europe-west2-a $(CLUSTER)

One day we'll do this all for you, but not today!

Common commands:

	- web-deployer publish:   Create a new deployable image of your application
	- web-deployer deploy:    Deploy a version of your application
`

var (
	verbose = false
)

type logger struct {
	*logrus.Logger
}

func (l *logger) Writer() *io.PipeWriter {
	return logger.WriterLevel(*l, logrus.DebugLevel)
}

func newRootCmd(args []string) *cobra.Command {
	logger := &logger{logrus.New()}

	cmd := &cobra.Command{
		Use:          "web-deployer",
		Short:        "Web Deployer for Kubernetes.",
		Long:         globalUsage,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.Formatter = &log.LogrusTextFormatter{}
			logger.Out = cmd.OutOrStdout()

			if verbose {
				logger.Level = logrus.DebugLevel
			}

			logger.Debugf("Using logger level %s...", logger.Level)
		},
	}

	cmd.AddCommand(
		newPublishCmd(logger),
		newDeployCmd(logger),
	)

	cmd.PersistentFlags().BoolVarP(&verbose, "verbose", "v", false, "verbose output")
	return cmd
}

func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
