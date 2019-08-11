package vkstreamapi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// AddRule func adds rules into stream
func AddRule(tag string, value string, key string) error {
	url := "https://streaming.vk.com/rules/?key=" + key

	client := http.Client{}

	var values Rules

	values.Rule.Value = value
	values.Rule.Tag = tag + "-=-" + value

	encodedRules, err := json.Marshal(values)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(encodedRules))
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	clientResp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	body, err := ioutil.ReadAll(clientResp.Body)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	log.Println(string(body))

	return nil
}
