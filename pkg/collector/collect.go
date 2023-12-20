package collector

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/spf13/cobra"
)

var configMapper = map[string]string{
	"kubeletAnonymousAuthArgumentSet":                  "authentication.anonymous.enabled",
	"kubeletAuthorizationModeArgumentSet":              "authorization.mode",
	"kubeletClientCaFileArgumentSet":                   "authentication.x509.clientCAFile",
	"kubeletReadOnlyPortArgumentSet":                   "readOnlyPort",
	"kubeletStreamingConnectionIdleTimeoutArgumentSet": "streamingConnectionIdleTimeout",
	"kubeletProtectKernelDefaultsArgumentSet":          "kernelMemcgNotification",
	"kubeletMakeIptablesUtilChainsArgumentSet":         "makeIPTablesUtilChains",
	"kubeletEventQpsArgumentSet":                       "eventRecordQPS",
	"kubeletRotateKubeletServerCertificateArgumentSet": "featureGates.RotateKubeletServerCertificate",
}

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
	log.SetFlags(log.LstdFlags | log.Lmicroseconds)
	cluster, err := GetCluster()
	if err != nil {
		return err
	}
	ctx, cancel := context.WithTimeout(cmd.Context(), time.Duration(10)*time.Minute)
	defer cancel()

	defer func() {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Println("Increase --timeout value")
		}
	}()
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
		nodeName := cmd.Flag("node").Value.String()
		configVal, err := getValuesFromkubeletConfig(ctx, nodeName, *cluster)
		if err != nil {
			return err
		}
		mergeConfigValues(nodeInfo, configVal)
		nodeData := Node{
			APIVersion: Version,
			Kind:       Kind,
			Type:       nodeType,
			Metadata:   map[string]string{"creationTimestamp": time.Now().Format(time.RFC3339)},
			Info:       nodeInfo,
		}
		outputFormat := cmd.Flag("output").Value.String()
		err = printOutput(nodeData, outputFormat, os.Stdout)
		if err != nil {
			return err
		}
	}
	return nil
}

func specByPlatfromVersion(platfrom string, version string) SpecVersion {
	return platfromSpec[fmt.Sprintf("%s-%s", platfrom, platfrom)]
}

func getValuesFromkubeletConfig(ctx context.Context, nodeName string, cluster Cluster) (map[string]*Info, error) {
	overrideConfig := make(map[string]*Info)
	data, err := cluster.clientSet.RESTClient().Get().AbsPath(fmt.Sprintf("/api/v1/nodes/%s/proxy/configz", nodeName)).DoRaw(ctx)
	if err != nil {
		return nil, err
	}
	nodeConfig := make(map[string]interface{})
	err = json.Unmarshal(data, &nodeConfig)
	if err != nil {
		return nil, err
	}
	values := nodeConfig["kubeletconfig"]
	for k, v := range configMapper {
		p := values
		var found bool
		splittedValues := StringToArray(v, ".")
		for _, sv := range splittedValues {
			next := p.(map[string]interface{})
			if k, ok := next[sv.(string)]; ok {
				found = true
				p = k
			}
		}
		if found {
			overrideConfig[k] = &Info{Values: []interface{}{p}}
		}
	}
	return overrideConfig, nil
}

func mergeConfigValues(configValues map[string]*Info, overrideConfig map[string]*Info) map[string]*Info {
	for k, v := range overrideConfig {
		configValues[k] = v
	}
	return configValues
}
