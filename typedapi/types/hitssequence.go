package types

import (
	"encoding/json"
)

type HitsSequence struct {
	Events []HitsEvent `json:"events"`

	JoinKeys []json.RawMessage `json:"join_keys,omitempty"`
}

func NewHitsSequence() *HitsSequence { _ = "STUB: not implemented"; return nil }
