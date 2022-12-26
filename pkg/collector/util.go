package collector

import (
	"strconv"
	"strings"
)

// StringToArray convert string with delimiter to array
func StringToArray(output string, delimiter string) []interface{} {
	if len(output) == 0 {
		return []interface{}{}
	}
	outputParts := strings.Split(output, delimiter)
	filterdParts := make([]interface{}, 0)
	for _, part := range outputParts {
		if intVal, err := strconv.Atoi(part); err == nil {
			filterdParts = append(filterdParts, intVal)
			continue
		}
		filterdParts = append(filterdParts, part)

	}
	return filterdParts
}

// SanitizeString snitize string from special characters
func SanitizeString(output string, replaceable map[string]string) string {
	for key, toReplace := range replaceable {
		output = strings.ReplaceAll(output, key, toReplace)
	}
	return output
}
