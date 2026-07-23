package types

import (
	"encoding/json"
)

type NodeInfoDiscover struct {
	NodeInfoDiscover map[string]json.RawMessage `json:"-"`
	SeedHosts        []string                   `json:"seed_hosts,omitempty"`
	SeedProviders    []string                   `json:"seed_providers,omitempty"`
	Type             *string                    `json:"type,omitempty"`
}

func (s *NodeInfoDiscover) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NodeInfoDiscover) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNodeInfoDiscover() *NodeInfoDiscover { _ = "STUB: not implemented"; return nil }
