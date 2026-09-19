package collector

import (
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	// Version resource version
	Version = "v1"
	// Kind resource kind
	Kind = "NodeInfo"
)

// LoadConfigParams load audit params data
func LoadConfigParams(nodeFileconfig string) (*Config, error) {
	var decodedNodeFileconfig []byte
	if nodeFileconfig == "" {
		return nil, fmt.Errorf("node file config is empty")
	}
	decodedNodeFileconfig, err := uncompressAndDecode(nodeFileconfig)
	if err != nil {
		fmt.Println("failed to read node file config")
		return nil, err
	}
	var np Config
	err = yaml.Unmarshal(decodedNodeFileconfig, &np)
	if err != nil {
		return nil, err
	}
	return &np, nil
}

func LoadKubeletMapping(kubletConfigMapping string) (map[string]string, error) {
	var fContent []byte
	var err error
	if kubletConfigMapping == "" {
		return nil, fmt.Errorf("kubletConfigMapping is empty")
	}
	fContent, err = uncompressAndDecode(kubletConfigMapping)
	if err != nil {
		fmt.Println("failed to read nodekubletConfigMapping")
		return nil, err
	}
	mapping := make(map[string]string)
	err = yaml.Unmarshal(fContent, &mapping)
	if err != nil {
		return nil, err
	}
	return mapping, nil
}

// SpecInfo spec info with require command to collect
type SpecInfo struct {
	Version  string    `yaml:"version"`
	Name     string    `yaml:"name"`
	Title    string    `yaml:"title"`
	Commands []Command `yaml:"commands"`
}

// Collector details of info to collect
type Command struct {
	Key      string `yaml:"key"`
	Title    string `yaml:"title"`
	Audit    string `yaml:"audit"`
	NodeType string `yaml:"nodeType"`
}

// Node output node data with info results
type Node struct {
	APIVersion string            `json:"apiVersion"`
	Kind       string            `json:"kind"`
	Metadata   map[string]string `json:"metadata"`
	Type       string            `json:"type"`
	Info       map[string]*Info  `json:"info"`
}

// Info command output result
type Info struct {
	Values interface{} `json:"values"`
}

type Config struct {
	Node NodeParams `yaml:"node"`
}
type Mapper struct {
	VersionMapping map[string][]SpecVersion `yaml:"version_mapping"`
}
type SpecVersion struct {
	Name           string
	Version        string `yaml:"cluster_version"`
	Op             string `yaml:"op"`
	CisSpecName    string `yaml:"spec_name"`
	CisSpecVersion string `yaml:"spec_version"`
}
type NodeParams struct {
	APIserver         Params            `yaml:"apiserver"`
	ControllerManager Params            `yaml:"controllermanager"`
	Scheduler         Params            `yaml:"scheduler"`
	Etcd              Params            `yaml:"etcd"`
	Proxy             Params            `yaml:"proxy"`
	KubeLet           Params            `yaml:"kubelet"`
	Flanneld          Params            `yaml:"flanneld"`
	VersionMapping    map[string]string `yaml:"version_mapping"`
}

type Params struct {
	Config            []string `yaml:"confs,omitempty"`
	DefaultConfig     string   `yaml:"defaultconf,omitempty"`
	KubeConfig        []string `yaml:"kubeconfig,omitempty"`
	DefaultKubeConfig string   `yaml:"defaultkubeconfig,omitempty"`
	DataDirs          []string `yaml:"datadirs,omitempty"`
	DefaultDataDir    string   `yaml:"defaultdatadir,omitempty"`
	Binaries          []string `yaml:"bins,omitempty"`
	DefaultBinaries   string   `yaml:"defaultbins,omitempty"`
	Services          []string `yaml:"svc,omitempty"`
	DefaultServices   string   `yaml:"defaultsvc,omitempty"`
	CAFile            []string `yaml:"cafile,omitempty"`
	DefaultCAFile     string   `yaml:"defaultcafile,omitempty"`
}
