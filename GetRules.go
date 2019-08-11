package vkstreamapi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// GetRules func returns rules from stream
func GetRules(tag string, value string, key string, endpoint string) error {
	url := "https://streaming.vk.com/rules/?key=" + key

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error: ", resp.Body)
	}


	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	log.Println(string(bodyBuf))

	return nil
}
