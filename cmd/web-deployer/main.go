package main

import (
	"os"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
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

func newRootCmd(args []string) *cobra.Command {
	logger := logrus.New()

	cmd := &cobra.Command{
		Use:          "web-deployer",
		Short:        "Web Deployer for Kubernetes.",
		Long:         globalUsage,
		SilenceUsage: true,
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			logger.Out = cmd.OutOrStdout()

			if verbose {
				logger.Info("Setting logger level to debug...")
				logger.Level = logrus.DebugLevel
			} else {
				logger.Info("Using logger level %s...", logger.Level)
			}
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
