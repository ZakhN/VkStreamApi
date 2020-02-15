package vkstreamapi

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetRules func returns rules from stream
func GetRules(key, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %v", resp.Body)
	}

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return bodyBuf, nil
}
