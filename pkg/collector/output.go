package collector

import (
	"encoding/json"
	"fmt"
	"io"
	"strconv"
	"strings"

	"github.com/olekukonko/tablewriter"
)

func printOutput(nodeData Node, output string, writer io.Writer) error {
	switch output {
	case "json":
		data, err := json.Marshal(nodeData)
		if err != nil {
			return err
		}
		fmt.Fprint(writer, string(data))
	case "table":
		data := make([][]string, 0)
		for key, ndata := range nodeData.Info {
			var results []string
			v, ok := ndata.Values.([]interface{})
			if !ok {
				return fmt.Errorf("type not supported")
			}
			for _, t := range v {
				switch n := t.(type) {
				case int:
					results = append(results, strconv.Itoa(n))
				case string:
					results = append(results, t.(string))
				}
			}
			if len(results) > 0 {
				joinedResults := join(results...)
				data = append(data, []string{key, joinedResults})
			}

		}
		table := tablewriter.NewWriter(writer)
		table.SetHeader([]string{"Key", "Value"})
		table.SetBorder(false) // Set Border to false
		table.AppendBulk(data) // Add Bulk Data
		table.Render()
	}
	return nil
}

func join(strs ...string) string {
	var sb strings.Builder
	for _, str := range strs {
		sb.WriteString(fmt.Sprintf(",%s", strings.TrimSpace(str)))
	}
	return strings.TrimPrefix(sb.String(), ",")
}
