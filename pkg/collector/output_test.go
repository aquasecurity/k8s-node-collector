package collector

import (
	"bytes"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPrintOutput(t *testing.T) {
	tests := []struct {
		name           string
		wantOutputFile string
		format         string
		nodeData       Node
	}{
		{
			name:           "print json format",
			format:         "json",
			wantOutputFile: "./testdata/fixture/output.json",
			nodeData: Node{
				APIVersion: Version,
				Type:       MasterNode,
				Kind:       "NodeInfo",
				Info: map[string]interface{}{
					"AdminConfFilePermissions":                 []interface{}{600},
					"CertificateAuthoritiesFilePermissions":    []interface{}{"root:root"},
					"ContainerNetworkInterfaceFilePermissions": []interface{}{700, 500},
				},
			},
		},
		{
			name:           "print table format",
			wantOutputFile: "./testdata/fixture/output.table",
			format:         "table",
			nodeData: Node{
				APIVersion: Version,
				Type:       MasterNode,
				Kind:       "NodeInfo",
				Info: map[string]interface{}{
					"ContainerNetworkInterfaceFilePermissions": []interface{}{700, 500},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			buff := bytes.NewBuffer([]byte{})
			printOutput(tt.nodeData, tt.format, buff)
			b, err := os.ReadFile(tt.wantOutputFile)
			assert.NoError(t, err)
			assert.Equal(t, buff.String(), string(b))
		})
	}
}
