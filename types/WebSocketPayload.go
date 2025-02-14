package types

import "encoding/json"

type WebSocketPayload struct {
	OP             int             `json:"op"`
	Data           json.RawMessage `json:"d"`
	SequenceNumber *int            `json:"s,omitempty"`
	EventName      string          `json:"t,omitempty"`
}
