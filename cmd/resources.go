package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"k8s.io/client-go/util/homedir"
)

type resourcesCmd struct {
	cmd *cobra.Command
}

func (r *resourcesCmd) execute() {
	if err := r.cmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func newResourcesCmd() *resourcesCmd {
	sync := &resourcesCmd{}

	data := &Data{}
	cmd := &cobra.Command{
		Use:           "resources",
		Aliases:       []string{"r"},
		Short:         "List resources for any kubernetes cluster",
		SilenceUsage:  true,
		SilenceErrors: true,
		Args:          cobra.ArbitraryArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("list kubernetes resources")

			return nil
		},
	}

	cmd.PersistentFlags().StringVarP(&data.Kubeconfig,
		"kubeconfig", "c", filepath.Join(homedir.HomeDir(), ".kube", "config"), "kubeconfig file to use for Kubernetes config")

	sync.cmd = cmd

	return sync
}
