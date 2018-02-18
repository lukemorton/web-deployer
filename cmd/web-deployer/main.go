package main

import (
	"os"
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

func newRootCmd(args []string) *cobra.Command {
	cmd := &cobra.Command{
		Use:          "web-deployer",
		Short:        "Web Deployer for Kubernetes.",
		Long:         globalUsage,
		SilenceUsage: true,
	}

	out := cmd.OutOrStdout()

	cmd.AddCommand(
		newPublishCmd(out),
	)

	return cmd
}

func main() {
	cmd := newRootCmd(os.Args[1:])
	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
