package collector

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestStringToArray(t *testing.T) {
	tests := []struct {
		name      string
		output    string
		delimiter string
		want      []interface{}
	}{
		{
			name:      "int results",
			output:    "600,700",
			delimiter: ",",
			want:      []interface{}{600, 700},
		},
		{
			name:      "string results",
			output:    "root:root",
			delimiter: ",",
			want:      []interface{}{"root:root"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := StringToArray(tt.output, tt.delimiter)
			assert.True(t, reflect.DeepEqual(got, tt.want))
		})
	}
}

func TestSanitizeString(t *testing.T) {
	tests := []struct {
		name        string
		output      string
		replaceable map[string]string
		want        string
	}{
		{
			name:        "repalce new line with comma",
			output:      "600\n700\n",
			replaceable: map[string]string{"\n": ","},
			want:        "600,700,",
		},
		{
			name:        "repalce regex with empty string",
			output:      "[^\"]\\S*'",
			replaceable: map[string]string{"[^\"]\\S*'": ""},
			want:        "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := SanitizeString(tt.output, tt.replaceable)
			assert.True(t, reflect.DeepEqual(got, tt.want))
		})
	}
}
