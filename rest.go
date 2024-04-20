package fugle_marketdata

import (
	"fmt"
	"net/http"
	"net/url"
)

const defaultRestClientEndpoint = "https://api.fugle.tw/marketdata/v1.0/stock"

// RestClientOption is a struct that represents the rest client option.
type RestClientOption struct {
	Endpoint string `json:"endpoint"`
	APIKey   string `json:"apiKey"`
}

// RestClient is a struct that represents the rest client.
type RestClient struct {
	option *RestClientOption

	httpclient *http.Client
}

// NewRestClient is a function used to create a new rest client.
func NewRestClient(option *RestClientOption) (*RestClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultRestClientEndpoint
	}

	_, err := url.ParseRequestURI(option.Endpoint)
	if err != nil {
		return nil, fmt.Errorf("parse rest client endpoint failed: %w", err)
	}

	if option.APIKey == "" {
		return nil, fmt.Errorf("rest client api key is required")
	}

	return &RestClient{
		option:     option,
		httpclient: http.DefaultClient,
	}, nil
}
