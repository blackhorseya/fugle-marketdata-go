package fugle_marketdata

// AuthEvent is a struct that represents the auth event response.
type AuthEvent struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
	} `json:"data"`
}

func (e *AuthEvent) GetEvent() string {
	return e.Event
}
