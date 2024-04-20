package fugle_marketdata

const defaultRestClientEndpoint = "https://api.fugle.tw/marketdata/v1.0/stock"

// RestClientOption is a struct that represents the rest client option.
type RestClientOption struct {
	Endpoint string `json:"endpoint"`
	APIKey   string `json:"apiKey"`
}

// RestClient is a struct that represents the rest client.
type RestClient struct {
	option *RestClientOption
}

// NewRestClient is a function used to create a new rest client.
func NewRestClient(option *RestClientOption) (*RestClient, error) {
	if option.Endpoint == "" {
		option.Endpoint = defaultRestClientEndpoint
	}

	// todo: 2024/4/20|sean|implement the rest client

	return &RestClient{
		option: option,
	}, nil
}
