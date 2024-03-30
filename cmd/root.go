package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

type Options struct {
	Kubeconfig string
}

type rootCmd struct {
	cmd *cobra.Command
}

func (r *rootCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newRootCmd() *rootCmd {

	root := &rootCmd{}
	cmd := &cobra.Command{
		Use:               "kagt",
		Short:             "Interact with Kubernetes resources",
		Long:              ``,
		SilenceUsage:      true,
		SilenceErrors:     true,
		Args:              cobra.NoArgs,
		ValidArgsFunction: cobra.NoFileCompletions,
	}

	cmd.AddCommand(
		newCheckCmd().cmd,
		newResourcesCmd().cmd,
	)
	root.cmd = cmd
	return root
}

func Execute() { newRootCmd().execute() }
