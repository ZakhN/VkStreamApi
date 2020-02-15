package vkstreamapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
)

// DeleteRule delete rule from stream
func DeleteRule(tag, url string) ([]byte, error) {
	req, err := http.NewRequest(http.MethodDelete, url, bytes.NewBuffer([]byte(tag)))
	if err != nil {
		return nil, fmt.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: ", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: ", resp.Body)
	}

	return body, nil
}
