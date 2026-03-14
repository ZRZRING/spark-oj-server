package utils

import (
	"strings"
)

// TODO: 后续可以多几种比较策略

func IsOutputEqual(actual, expected string) bool {
	return strings.TrimSpace(actual) == strings.TrimSpace(expected)
}
