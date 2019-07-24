package VkStreamApi

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func GetRules(key string) error {
	url := fmt.Sprintf("")

	req, err := http.NewRequest("GET", url, nil)

	client := &http.Client{}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	defer resp.Body.Close()

	bodyBuf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error: ", err)
	}

	log.Println(string(bodyBuf))

	return nil
}