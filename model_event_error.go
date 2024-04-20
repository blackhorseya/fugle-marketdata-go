package fugle_marketdata

// ErrorEventResponse is a struct that represents the error event response.
type ErrorEventResponse struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
	} `json:"data"`
}

func (e *ErrorEventResponse) GetEvent() string {
	return e.Event
}
