package fugle_marketdata

// ErrorEvent is a struct that represents the error event response.
type ErrorEvent struct {
	Event string `json:"event"`
	Data  struct {
		Message string `json:"message"`
	} `json:"data"`
}

func (e *ErrorEvent) GetEvent() string {
	return e.Event
}
