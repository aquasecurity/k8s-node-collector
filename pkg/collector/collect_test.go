package collector

import (
	"encoding/json"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestParseNodeConfig(t *testing.T) {
	tests := []struct {
		name                   string
		nodeConfigFile         string
		expextedNodeConfigFile map[string]*Info
	}{
		{
			name:           "parse node config",
			nodeConfigFile: "./testdata/fixture/node_config.json",
			expextedNodeConfigFile: map[string]*Info{
				"kubeletAnonymousAuthArgumentSet": {
					Values: []interface{}{"false"},
				},
				"kubeletAuthorizationModeArgumentSet": {
					Values: []interface{}{"Webhook"},
				},
				"kubeletClientCaFileArgumentSet": {
					Values: []interface{}{"/etc/kubernetes/pki/ca.crt"},
				},
				"kubeletEventQpsArgumentSet": {
					Values: []interface{}{5.0},
				},
				"kubeletMakeIptablesUtilChainsArgumentSet": {
					Values: []interface{}{"true"},
				},
				"kubeletStreamingConnectionIdleTimeoutArgumentSet": {
					Values: []interface{}{"4h0m0s"},
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
			m, err := getValuesFromkubeletConfig(nodeConfig)
			assert.NoError(t, err)
			for k, v := range m {
				assert.Equal(t, v, tt.expextedNodeConfigFile[k])
			}
		})
	}
}
