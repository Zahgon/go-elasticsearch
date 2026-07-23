package types

import (
	"encoding/json"
)

type NodeUsage struct {
	Aggregations map[string]json.RawMessage `json:"aggregations"`

	RestActions map[string]int `json:"rest_actions"`

	Since int64 `json:"since"`

	Timestamp int64 `json:"timestamp"`
}

func (s *NodeUsage) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeUsage() *NodeUsage { _ = "STUB: not implemented"; return nil }
