package collector

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

// CollectNodeData run spec audit command and output it result data
func CollectNodeData(cmd *cobra.Command) error {
	specVersion := cmd.Flag("spec").Value.String()
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
		nodeInfo := make(map[string]*Info)
		if infoCollector.Version != specVersion {
			continue
		}
		for _, ci := range infoCollector.Collectors {
			if ci.NodeType != nodeType && nodeType != MasterNode {
				continue
			}
			output, err := shellCmd.Execute(ci.Audit)
			if err != nil {
				fmt.Print(err)
			}
			values := StringToArray(output, ",")
			nodeInfo[ci.Key] = &Info{Values: values}
		}
		nodeData := Node{
			APIVersion: Version,
			Kind:       Kind,
			Type:       nodeType,
			Metadata:   map[string]string{"creationTimestamp": time.Now().Format(time.RFC3339)},
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
