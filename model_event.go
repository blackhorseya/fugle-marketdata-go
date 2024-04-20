package fugle_marketdata

import (
	"encoding/json"
)

// IEvent is an interface that represents the event.
type IEvent interface {
	// GetEvent is a function used to get the event.
	GetEvent() string
}

// UnknownEvent is a struct that represents the unknown event.
type UnknownEvent struct {
	Event string          `json:"event"`
	Data  json.RawMessage `json:"data"`
}

func (e *UnknownEvent) GetEvent() string {
	return e.Event
}

// UnmarshalEvent is a function used to unmarshal the event.
func UnmarshalEvent(data []byte) (IEvent, error) {
	var raw map[string]json.RawMessage
	err := json.Unmarshal(data, &raw)
	if err != nil {
		return nil, err
	}

	var event string
	err = json.Unmarshal(raw["event"], &event)
	if err != nil {
		return nil, err
	}

	switch event {
	case "error":
		var e *ErrorEvent
		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		return e, nil
	default:
		var e *UnknownEvent
		err = json.Unmarshal(data, &e)
		if err != nil {
			return nil, err
		}

		return e, nil
	}
}
