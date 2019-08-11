package vkstreamapi

// VkAPIResponse struct
type VkAPIResponse struct {
	Response Response `json:"response"`
}

type Response struct {
	Endpoint string `json:"endpoint"`
	Key      string `json:"key"`
}

// Rules struct
type Rules struct {
	Rule `json:"rule"`
}

type Rule struct {
	Value string `json:"value"`
	Tag   string `json:"tag"`
}
