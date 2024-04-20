package fugle_marketdata

// AuthEventRequest is a struct that represents the auth event request.
type AuthEventRequest struct {
	Event string `json:"event"`
}

func (e *AuthEventRequest) GetEvent() string {
	return e.Event
}

// AuthEventResponse is a struct that represents the auth event response.
type AuthEventResponse struct {
	Event string `json:"event"`
}

func (e *AuthEventResponse) GetEvent() string {
	return e.Event
}
