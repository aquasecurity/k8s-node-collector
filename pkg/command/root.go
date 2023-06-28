package cmd

import (
	"fmt"
	"os"

	"github.com/aquasecurity/k8s-node-collector/pkg/collector"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "json", "Output format. One of table|json")
	rootCmd.PersistentFlags().StringP("spec", "s", "cis", " spec name.  default: cis")
	rootCmd.PersistentFlags().StringP("version", "v", "1.23", "spec version.  default: 1.23")
}

var rootCmd = &cobra.Command{
	Use:   "node-collector",
	Short: "k8s-node-collector extract file system info from cluster Node",
	Long:  `A tool which provide a way to extract k8s info which is not accessible via apiserver from node cluster based on pre-define commands`,
	RunE: func() func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			return collector.CollectNodeData(cmd)
		}
	}(),
}

// Execute CLI commands
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
