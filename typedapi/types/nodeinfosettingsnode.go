package types

import (
	"encoding/json"
)

type NodeInfoSettingsNode struct {
	Attr                 map[string]json.RawMessage `json:"attr"`
	MaxLocalStorageNodes *string                    `json:"max_local_storage_nodes,omitempty"`
	Name                 string                     `json:"name"`
}

func (s *NodeInfoSettingsNode) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeInfoSettingsNode() *NodeInfoSettingsNode { _ = "STUB: not implemented"; return nil }
