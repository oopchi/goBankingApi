package util

import "strings"

var statusCode = map[string]int{
	"active":   1,
	"inactive": 0,
}

var statusText = map[int]string{
	1: "active",
	0: "inactive",
}

func StatusCode(status string) int {
	if val, ok := statusCode[strings.Trim(strings.ToLower(status), " ")]; ok {
		return val
	}

	return -1
}

func StatusText(status int) string {
	if val, ok := statusText[status]; ok {
		return val
	}

	return "inactive"
}
