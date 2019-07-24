package VkStreamApi

type vkAPIResponse struct {
	Response struct {
		Endpoint string `json:"endpoint"`
		Key      string `json:"key"`
	}
}

// Rules struct
type Rules struct {
	Rule struct {
		Value string `json:"value"`
		Tag   string `json:"tag"`
	} `json:"rule"`
}
