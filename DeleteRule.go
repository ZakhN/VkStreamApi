package vkstreamapi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// DeleteRule delete rule from stream
func DeleteRule(tag string, value string, key string) error {
	url := "https://streaming.vk.com/rules/?key="+key

	tag = fmt.Sprintf(`{"tag":"%s"}`, tag+"-=-"+value)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte(tag)))
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	defer resp.Body.Close()
	
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	if resp.StatusCode != 200 {
		return fmt.Errorf("error: ", resp.Body)
	}

	log.Println(string(body))

	return nil
}
