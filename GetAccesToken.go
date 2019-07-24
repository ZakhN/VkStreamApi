package VkStreamApi

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

func GetAccesToken() {
	resp, err := http.Get("")
	if err != nil {
		log.Fatal("firehose api authorization failed:", err)
	}
	bodyBuf, err := ioutil.ReadAll(resp.Body)
	var v vkAPIResponse

	if err := json.Unmarshal(bodyBuf, &v); err != nil {
		log.Fatal("unmarshal response json failed:", err)
	}
	log.Println("host:", v.Response.Endpoint, "key:", v.Response.Key)
}
