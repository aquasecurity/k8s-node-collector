package collector

import (
	"fmt"
	"os"
	"time"

	"github.com/spf13/cobra"
)

type SpecVersion struct {
	Name    string
	Version string
}

var platfromSpec = map[string]SpecVersion{
	"k8s-1.23": {
		Name:    "cis",
		Version: "1.23",
	},
}

// CollectData run spec audit command and output it result data
func CollectData(cmd *cobra.Command, target string) error {
	cluster, err := GetCluster()
	if err != nil {
		return err
	}
	p, err := cluster.Platfrom()
	if err != nil {
		return err
	}
	shellCmd := NewShellCmd()
	nodeType, err := shellCmd.FindNodeType()
	if err != nil {
		return err
	}
	infoCollectorMap, err := LoadConfig(target)
	if err != nil {
		return err
	}
	specName := cmd.Flag("spec").Value.String()
	specVersion := cmd.Flag("version").Value.String()
	sv := SpecVersion{Name: specName, Version: specVersion}
	if len(sv.Name) == 0 || len(sv.Version) == 0 {
		sv = specByPlatfromVersion(p.Name, p.Version)
	}
	for _, infoCollector := range infoCollectorMap {
		nodeInfo := make(map[string]*Info)
		if !(infoCollector.Version == sv.Version && infoCollector.Name == sv.Name) {
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

func specByPlatfromVersion(platfrom string, version string) SpecVersion {
	return platfromSpec[fmt.Sprintf("%s-%s", platfrom, platfrom)]
}
