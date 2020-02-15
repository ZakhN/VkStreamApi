package vkstreamapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// AddRule func adds rules into stream
func AddRule(tag, value, url string) ([]byte, error) {
	values := Rules{
		Rule: Rule{
			Value: value,
			Tag:   tag,
		},
	}

	encodedRules, err := json.Marshal(values)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(encodedRules))
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}
	resp, err := client.Do(req)
	defer resp.Body.Close()

	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %v", resp.Body)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return body, nil
}
