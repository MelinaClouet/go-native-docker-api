package services

import "fmt"

func mapToSlice(env map[string]string) []string {
	var result []string
	for k, v := range env {
		if k != "" && v != "" {
			result = append(result, fmt.Sprintf("%s=%s", k, v))
		}
	}
	return result
}
