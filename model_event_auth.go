package fugle_marketdata

// AuthEventRequest is a struct that represents the auth event request.
type AuthEventRequest struct {
	Event string `json:"event"`
	Data  struct {
		APIKey string `json:"apikey"`
	} `json:"data"`
}

func (e *AuthEventRequest) GetEvent() string {
	return e.Event
}

// AuthEventResponse is a struct that represents the auth event response.
type AuthEventResponse struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
	} `json:"data"`
}

func (e *AuthEventResponse) GetEvent() string {
	return e.Event
}
