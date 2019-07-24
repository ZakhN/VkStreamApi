package VkStreamApi

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func DeleteRule(tag string, value string, key string, endpoint string) error {
	url := ""

	tag := fmt.Sprintf(`{"tag":"%s"}`, tag+"./."+value)

	req, err := http.NewRequest("DELETE", url, bytes.NewBuffer([]byte(tag)))
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	req.Header.Set("Content-Type", "application/json")

	client := http.Client{}

	resp, err := client.Do(req)
	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	log.Println(tag, "***RULES DELETED***")
	log.Println(string(bodyBuf))

	return nil
}
