package ssllabs

import (
	"fmt"
	"log"
	"net/http"
	"time"
  "io/ioutil"
  "encoding/json"
)

func callSslLabs(domain string) (bool, interface{}, error) {
  url := fmt.Sprintf("https://api.ssllabs.com/api/v3/analyze?host=%s", domain)
	fmt.Println(url)

	spaceClient := http.Client{
		Timeout: time.Second * 6, // Maximum of 2 secs
	}

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return false, nil, err
	}

	req.Header.Set("User-Agent", "truora-test-fulit103")

  res, getErr := spaceClient.Do(req)
	if getErr != nil {
		log.Println(getErr)
    return false, nil, getErr
	}

  body, readErr := ioutil.ReadAll(res.Body)
	if readErr != nil {
		log.Println(readErr)
    return false, nil, readErr
	}

  var f interface{}
  errJSON := json.Unmarshal(body, &f)

  if errJSON != nil {
		log.Println(errJSON)
    return false, nil, errJSON
	}

  m := f.(map[string]interface{})
  fmt.Println("---------")
  fmt.Println(m)
  fmt.Println("---------")
  fmt.Println("#########")
  fmt.Println(m["status"])
  fmt.Println("#########")

  for k, v := range m {
      switch vv := v.(type) {
      case string:
          fmt.Println(k, "is string", vv)
      case float64:
          fmt.Println(k, "is float64", vv)
      case []interface{}:
          fmt.Println(k, "is an array:")
          for i, u := range vv {
              fmt.Println(i, u)
          }
      default:
          fmt.Println(k, "is of a type I don't know how to handle")
      }
  }

  if m["status"] == "READY" {
    return true, m, nil
  }

  return false, m, nil
}

// ScannDomain Escanea el dominio pasado con ssllibas
func ScannDomain(domain string) {

  for {
    fmt.Println("##### CallSslLabs #####")
    ready, data, error := callSslLabs(domain)

    if error != nil {
      log.Println(error)
    }
    if data != nil {
      fmt.Println("Ready")
    }

    if ready == true {
      break
    }
    time.Sleep(time.Second * 10)
  }

}
