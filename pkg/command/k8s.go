package cmd

import (
	"github.com/aquasecurity/k8s-node-collector/pkg/collector"
	"github.com/spf13/cobra"
)

const (
	subCommandK8s = "k8s"
)

func init() {
	rootCmd.AddCommand(k8sCmd)
}

var k8sCmd = &cobra.Command{
	Use:   subCommandK8s,
	Short: "k8s-node-collector extract file system info from cluster Node",
	Long:  `A tool which provide a way to extract k8s info which is not accessible via apiserver from node cluster based on pre-define commands`,
	RunE: func() func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			return collector.CollectData(cmd, subCommandK8s)
		}
	}(),
}
