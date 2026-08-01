package lambda

import (
	"encoding/json"
)

type Event struct {
	Collection string          `json:"collection"`
	Command    string          `json:"command"`
	Filter     json.RawMessage `json:"filter"`
	Data       json.RawMessage `json:"data"`
}
