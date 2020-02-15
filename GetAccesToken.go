package vkstreamapi

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

func GetAccessToken(url string) (*VkAPIResponse, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var response VkAPIResponse

	if err := json.Unmarshal(bodyBuf, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
