package utils

import "strings"

func FormatTemplate(template string, replacements map[string]string) string {
	for placeholder, replacement := range replacements {
		template = strings.ReplaceAll(template, placeholder, replacement)
	}
	return template
}
