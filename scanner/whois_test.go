package scanner

import (
	"testing"

	whois "github.com/likexian/whois-go"
)

func TestWhoisOwner(t *testing.T) {
	result, err := whois.Whois("54.88.139.186")
	if err != nil {
		t.Error("no coincide")
	}

	owner, err := getWhoisField(result, "owner")
	if err != nil {
		t.Error("No Encontro Owner")
	}
	if owner == "Amazon Technologies Inc." {
		println("Owner: ", owner)
	} else {
		t.Error("No Encontro Owner")
		println(result)
	}
}

func TestWhoisCountry(t *testing.T) {
	result, err := whois.Whois("54.88.139.186")
	if err != nil {
		t.Error("no coincide")
	}

	country, err := getWhoisField(result, "country")
	if err != nil {
		t.Error("No Encontro Country")
	}
	if country == "US" {
		println("Owner: ", country)
	} else {
		t.Error("No Encontro Owner")
		println(result)
	}
}
