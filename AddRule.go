package VkStreamApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func AddRule(tag string, value string) error {
	url := ""

	client := http.Client{}

	var values Rules

	values.Rule.Value = value
	values.Rule.Tag = tag + "./." + value

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

	return nil
}
