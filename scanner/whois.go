package scanner

import (
	"errors"
	"strings"
)

func getWhoisField(whois string, field string) (string, error) {
	array := strings.Split(whois, "\n")
	for _, line := range array {
		if strings.Contains(line, field) {
			return strings.TrimSpace(strings.Split(line, ":")[1]), nil
		}
	}
	return "", errors.New("Not found")
}
