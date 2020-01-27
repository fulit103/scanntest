package scanner

import (
	"errors"
	"strings"
)

func getWhoisField(whois string, field string) (string, error) {
	array := strings.Split(whois, "\n")
	for _, line := range array {
		lineTmp := strings.ToLower(line)
		if field == "owner" {
			if strings.Contains(lineTmp, strings.ToLower("OrgName:")) || strings.Contains(lineTmp, "owner:") {
				return strings.TrimSpace(strings.Split(line, ":")[1]), nil
			}
		}
		if field == "country" {
			if strings.Contains(lineTmp, "country:") {
				return strings.TrimSpace(strings.Split(line, ":")[1]), nil
			}
		}
	}
	return "", errors.New("Not found")
}
