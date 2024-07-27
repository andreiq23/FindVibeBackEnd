package utils

import (
	"errors"
	"strings"
)

func CleanString(s string) (string, error) {
	if s == "" {
		return "", errors.New("empty string")
	}
	return strings.Replace(s, " ", "+", -1), nil
}
