package scanner

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

// CallSslLabs call api ssllabs
func CallSslLabs(domain string) (bool, map[string]interface{}, interface{}, error) {
	url := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", domain)
	fmt.Println(url)

	spaceClient := http.Client{
		Timeout: time.Second * 6, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, nil, nil, err
	}

	req.Header.Set("User-Agent", "truora-test-fulit103")

	res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Println(getErr)
		return false, nil, nil, getErr
	}

	body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
		return false, nil, nil, readErr
	}

	var f interface{}
	errJSON := json.Unmarshal(body, &f)

	if errJSON != nil {
		log.Println(errJSON)
		return false, nil, nil, errJSON
	}

	m := f.(map[string]interface{})

	status := m["status"]

	if status == "READY" {
		return true, m, status, nil
	}

	return false, m, status, nil
}
