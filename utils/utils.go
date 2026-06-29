package utils

import "fmt"

func MapToSlice(env map[string]string) []string {
	var result []string
	for k, v := range env {
		if k != "" && v != "" {
			result = append(result, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return result
}
