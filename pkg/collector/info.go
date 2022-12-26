package collector

import (
	"embed"
	"fmt"

	"gopkg.in/yaml.v3"
)

const (
	configFolder = "config"
	// Version resource version
	Version = "v1"
	// Kind resource kind
	Kind = "NodeInfo"
)

//go:embed config
var config embed.FS

// LoadConfig load audit commands specification from config file
func LoadConfig() (map[string]*SpecInfo, error) {
	dirEntries, err := config.ReadDir(configFolder)
	if err != nil {
		return nil, err
	}
	specInfoMap := make(map[string]*SpecInfo)
	for _, entry := range dirEntries {
		fContent, err := config.ReadFile(fmt.Sprintf("%s/%s", configFolder, entry.Name()))
		if err != nil {
			return nil, err
		}
		si, err := getSpecInfo(string(fContent))
		if err != nil {
			return nil, err
		}
		specInfoMap[si.Name] = si
	}
	return specInfoMap, nil
}

// SpecInfo spec info with require comand to collect
type SpecInfo struct {
	Version    string      `yaml:"version"`
	Name       string      `yaml:"name"`
	Title      string      `yaml:"title"`
	Collectors []Collector `yaml:"collectors"`
}

// Collector details of info to collect
type Collector struct {
	Key      string `yaml:"key"`
	Title    string `yaml:"title"`
	Audit    string `yaml:"audit"`
	NodeType string `yaml:"nodeType"`
}

func getSpecInfo(info string) (*SpecInfo, error) {
	var specInfo SpecInfo
	err := yaml.Unmarshal([]byte(info), &specInfo)
	if err != nil {
		return nil, err
	}
	return &specInfo, nil
}

// Node output node data with info results
type Node struct {
	APIVersion string                 `json:"apiVersion"`
	Kind       string                 `json:"kind"`
	Type       string                 `json:"type"`
	Info       map[string]interface{} `json:"info"`
}
