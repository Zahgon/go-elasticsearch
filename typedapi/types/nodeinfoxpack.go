package types

import (
	"encoding/json"
)

type NodeInfoXpack struct {
	License      *NodeInfoXpackLicense      `json:"license,omitempty"`
	Ml           *NodeInfoXpackMl           `json:"ml,omitempty"`
	Notification map[string]json.RawMessage `json:"notification,omitempty"`
	Security     NodeInfoXpackSecurity      `json:"security"`
}

func NewNodeInfoXpack() *NodeInfoXpack { _ = "STUB: not implemented"; return nil }
