package collector

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"os"
	"reflect"
	"strings"
	"testing"

	"github.com/dsnet/compress/bzip2"
	"github.com/stretchr/testify/assert"
	"sigs.k8s.io/kustomize/kyaml/yaml"
)

func TestParseNodeConfig(t *testing.T) {
	tests := []struct {
		name                   string
		nodeConfigFile         string
		mappingFile            string
		expectedNodeConfigFile map[string]*Info
	}{
		{
			name:           "parse node config",
			nodeConfigFile: "./testdata/fixture/node_config.json",
			mappingFile:    "./testdata/fixture/kubeletconfig-mapping.yaml",
			expectedNodeConfigFile: map[string]*Info{
				"kubeletAnonymousAuthArgumentSet": {
					Values: []interface{}{"false"},
				},
				"kubeletAuthorizationModeArgumentSet": {
					Values: []interface{}{"Webhook"},
				},
				"kubeletClientCaFileArgumentSet": {
					Values: []interface{}{"/etc/kubernetes/certs/ca.crt"},
				},
				"kubeletEventQpsArgumentSet": {
					Values: []interface{}{0.0},
				},
				"kubeletMakeIptablesUtilChainsArgumentSet": {
					Values: []interface{}{"true"},
				},
				"kubeletStreamingConnectionIdleTimeoutArgumentSet": {
					Values: []interface{}{"4h0m0s"},
				},
				"kubeletOnlyUseStrongCryptographic": {
					Values: []interface{}{"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256",
						"TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256",
						"TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305",
						"TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384",
						"TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305",
						"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384",
						"TLS_RSA_WITH_AES_256_GCM_SHA384",
						"TLS_RSA_WITH_AES_128_GCM_SHA256"},
				},
			}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			data, err := os.ReadFile(tt.nodeConfigFile)
			assert.NoError(t, err)
			nodeConfig := make(map[string]interface{})
			err = json.Unmarshal(data, &nodeConfig)
			assert.NoError(t, err)
			mf, err := os.ReadFile(tt.mappingFile)
			assert.NoError(t, err)
			bzip2CompressData, err := bzip2Compress(mf)
			assert.NoError(t, err)
			encodedMapping := base64.StdEncoding.EncodeToString(bzip2CompressData)
			mapping, err := LoadKubeletMapping(encodedMapping)
			assert.NoError(t, err)
			m := getValuesFromkubeletConfig(nodeConfig, mapping)
			for k, v := range m {
				if _, ok := tt.expectedNodeConfigFile[k]; ok {
					assert.Equal(t, v, tt.expectedNodeConfigFile[k])
				}
			}
		})
	}
}

func bzip2Compress(data []byte) ([]byte, error) {
	var buf bytes.Buffer
	w, err := bzip2.NewWriter(&buf, &bzip2.WriterConfig{Level: bzip2.DefaultCompression})
	if err != nil {
		return []byte{}, err
	}

	_, err = w.Write(data)
	if err != nil {
		return []byte{}, err
	}
	err = w.Close()
	if err != nil {
		return []byte{}, err
	}
	return buf.Bytes(), nil
}

func TestSpecByVersionName(t *testing.T) {
	tests := []struct {
		name               string
		versionMappingfile string
		platfrom           Platform
		wantSpec           string
	}{
		{
			name:               "k8s cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "k8s", Version: "1.21"},
			wantSpec:           "k8s-cis-1.23.0",
		},
		{
			name:               "aks cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "aks", Version: "1.21"},
			wantSpec:           "aks-cis-1.0.0",
		},
		{
			name:               "eks cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "eks", Version: "1.21"},
			wantSpec:           "eks-cis-1.2.0",
		},
		{
			name:               "gke cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "gke", Version: "1.21"},
			wantSpec:           "gke-cis-1.2.0",
		},
		{
			name:               "rke2 cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "rke2", Version: "1.21"},
			wantSpec:           "rke2-cis-1.24.0",
		},
		{
			name:               "ocp cis spec",
			versionMappingfile: "./testdata/fixture/mapping.yaml",
			platfrom:           Platform{Name: "ocp", Version: "4.0"},
			wantSpec:           "rh-cis-1.0",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f, err := os.ReadFile(tt.versionMappingfile)
			assert.NoError(t, err)
			var mapper Mapper
			err = yaml.Unmarshal(f, &mapper)
			assert.NoError(t, err)
			gotSpec := specByPlatfromVersion(tt.platfrom, mapper.VersionMapping)
			assert.Equal(t, gotSpec, tt.wantSpec)
		})
	}
}

func TestPlatfromVersion(t *testing.T) {
	tests := []struct {
		name    string
		version string
		want    string
	}{
		{
			name:    "k8s version",
			version: "v1.23.2",
			want:    "1.23",
		},
		{
			name:    "eks version",
			version: "v1.23.17-eks-8ccc7ba",
			want:    "1.23",
		},
		{
			name:    "aks version",
			version: "v1.23.17-eks-8ccc7ba",
			want:    "1.23",
		},
		{
			name:    "gke version",
			version: "v1.23.10-gke.2300",
			want:    "1.23",
		},
		{
			name:    "rke2 version",
			version: "v1.23.11+rke2r1",
			want:    "1.23",
		},
		{
			name:    "ocp version",
			version: "v1.23.15+c763d11",
			want:    "1.23",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getPlatformInfoFromVersion(tt.version)
			assert.Equal(t, got.Version, tt.want)
		})
	}
}

func TestNodeCommamnd(t *testing.T) {
	tests := []struct {
		name             string
		commandsFilePath string
		want             []Command
	}{
		{
			name:             "k8s version",
			commandsFilePath: "./testdata/fixture/single-check.yaml",
			want: []Command{
				{
					Key:      "kubeAPIServerSpecFilePermission",
					Title:    "API server pod specification file permissions",
					NodeType: "master",
					Audit:    "stat -c %a $apiserver.confs",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fd, err := os.ReadFile(tt.commandsFilePath)
			assert.NoError(t, err)
			commands, err := CompressAndEncode(fd)
			assert.NoError(t, err)
			got, err := GetNodesCommands(string(commands), map[string]string{}, "master")
			assert.NoError(t, err)
			assert.True(t, reflect.DeepEqual(got, tt.want))
		})
	}
}

func CompressAndEncode(data []byte) (string, error) {
	cm, err := bzip2Compress(data)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(cm), nil
}

// fakeShell replays a recorded process list the way Shell.Execute returns it,
// with the lines joined by commas.
type fakeShell struct {
	procList []string
}

func (f *fakeShell) Execute(string) (string, error) {
	return strings.Join(f.procList, ","), nil
}

func (f *fakeShell) FindNodeType() (string, error) {
	return MasterNode, nil
}

func TestBinLookup(t *testing.T) {
	// the binaries the shipped config looks for
	proxyBins := []string{"kube-proxy", "hyperkube proxy", "hyperkube kube-proxy", "proxy", "openshift start network"}
	tests := []struct {
		name     string
		bins     []string
		procList []string
		want     string
	}{
		{
			name: "kube-proxy is running",
			bins: proxyBins,
			procList: []string{
				"1 /sbin/init",
				"842 /usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf --hostname-override=node1",
			},
			want: "kube-proxy",
		},
		{
			name: "sub command is matched as its own argument",
			bins: proxyBins,
			procList: []string{
				"1 /sbin/init",
				"842 /usr/bin/hyperkube proxy --config=/var/lib/kube-proxy/config.conf",
			},
			want: "hyperkube",
		},
		{
			name: "a different sub command is not a match",
			bins: proxyBins,
			procList: []string{
				"1 /sbin/init",
				"842 /usr/bin/hyperkube kubelet --config=/var/lib/kubelet/config.yaml",
			},
			want: "kube-proxy",
		},
		{
			// https://github.com/aquasecurity/k8s-node-collector/issues/254
			name: "the process title of an unrelated process is not a match",
			bins: proxyBins,
			procList: []string{
				"1 /sbin/init",
				"318 dhcpcd: [privileged proxy] ens18 [ip4] [ip6]",
			},
			want: "kube-proxy",
		},
		{
			name: "an unrelated binary ending in the searched name is not a match",
			bins: proxyBins,
			procList: []string{
				"1 /sbin/init",
				"624 /usr/local/sbin/haproxy -f /etc/haproxy/haproxy.cfg",
			},
			want: "kube-proxy",
		},
		{
			name: "a command line holding a comma does not break the following record",
			bins: proxyBins,
			procList: []string{
				"1 /usr/local/bin/kube-apiserver --enable-admission-plugins=NodeRestriction,PodSecurity",
				"842 /usr/local/bin/kube-proxy --config=/var/lib/kube-proxy/config.conf",
			},
			want: "kube-proxy",
		},
		{
			name:     "no process is running any of the binaries",
			bins:     proxyBins,
			procList: []string{"1 /sbin/init"},
			want:     "kube-proxy",
		},
		{
			name:     "no binaries configured",
			bins:     []string{},
			procList: []string{"1 /sbin/init"},
			want:     "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := binLookup(tt.bins, "kube-proxy", &fakeShell{procList: tt.procList})
			assert.Equal(t, tt.want, got)
		})
	}
}

func TestLoadNodeCommands(t *testing.T) {
	const spec = `---
version: "1.0"
name: k8s-cis
title: CIS Kubernetes Benchmarks
commands:
  - key: kubeProxyKubeconfigFileOwnership
    title: kube-proxy kubeconfig file ownership
    nodeType: worker
    audit: output=` + "`" + `stat -c %U:%G $(ps -ef | grep $proxy.bins | grep 'kubeconfig' | head -n1)` + "`" + `
`
	tests := []struct {
		name         string
		configMap    map[string]string
		wantErr      string
		wantCommands int
	}{
		{
			name:         "spec is substituted and parsed",
			configMap:    map[string]string{"$proxy.bins": "kube-proxy"},
			wantCommands: 1,
		},
		{
			// https://github.com/aquasecurity/k8s-node-collector/issues/154
			name:      "a broken substitution is reported as the parse error it is",
			configMap: map[string]string{"$proxy.bins": "dhcpcd:"},
			wantErr:   "mapping values are not allowed",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			encoded, err := CompressAndEncode([]byte(spec))
			assert.NoError(t, err)
			commands, err := loadNodeCommands(encoded, tt.configMap, "worker")
			if tt.wantErr != "" {
				assert.ErrorContains(t, err, tt.wantErr)
				return
			}
			assert.NoError(t, err)
			assert.Len(t, commands, tt.wantCommands)
		})
	}
}
