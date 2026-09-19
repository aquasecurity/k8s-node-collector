package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.PersistentFlags().StringP("output", "o", "json", "Output format. One of table|json")
	rootCmd.PersistentFlags().StringP("spec-name", "s", "", "spec name. example: k8s-cis")
	rootCmd.PersistentFlags().StringP("spec-version", "v", "", "spec version. example 1.23.0")
	rootCmd.PersistentFlags().StringP("cluster-version", "c", "", "cluster version. example 1.23.0")
	rootCmd.PersistentFlags().StringP("node", "n", "", "node name")
	rootCmd.PersistentFlags().StringP("kubelet-config", "", "", "kubelet config via api /api/v1/nodes/<>/proxy/configz encoded to base64")
	rootCmd.PersistentFlags().StringP("spec-version-mapping", "", "", "k8s spec-version mapping encoded to base64")
	rootCmd.PersistentFlags().StringP("node-config", "", "", "k8s node file config encoded to base64")
	rootCmd.PersistentFlags().StringP("node-commands", "", "", "k8s node commands to be executed encoded to base64")
	rootCmd.PersistentFlags().StringP("kubelet-config-mapping", "", "", "kubelet config api mapping encoded to base64")
}

var rootCmd = &cobra.Command{
	Use:     "node-collector",
	Example: "node-collector k8s [flags]",
	Short:   "trivy-collector extract file system info",
	Long:    `A tool which provides a way to extract file info which is not accessible via pre-defined commands`,
	RunE: func() func(cmd *cobra.Command, args []string) error {
		return func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
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
