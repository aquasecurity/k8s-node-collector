package collector

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// CollectNodeData run spec audit command and output it result data
func CollectNodeData(cmd *cobra.Command) error {
	shellCmd := NewShellCmd()
	nodeType, err := shellCmd.FindNodeType()
	if err != nil {
		return err
	}
	infoCollectorMap, err := LoadConfig()
	if err != nil {
		return err
	}
	for _, infoCollector := range infoCollectorMap {
		nodeInfo := make(map[string]interface{})
		for _, ci := range infoCollector.Collectors {
			if ci.NodeType != nodeType && nodeType != MasterNode {
				continue
			}
			output, err := shellCmd.Execute(ci.Audit)
			if err != nil {
				fmt.Print(err)
			}
			values := StringToArray(output, ",")
			nodeInfo[ci.Key] = values
		}
		nodeData := Node{
			APIVersion: Version,
			Kind:       Kind,
			Type:       nodeType,
			Info:       nodeInfo,
		}
		outputFormat := cmd.Flag("output").Value.String()
		err := printOutput(nodeData, outputFormat, os.Stdout)
		if err != nil {
			return err
		}
	}
	return nil
}
