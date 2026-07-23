package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardroutingstate"
)

type ShardRouting struct {
	Node           string                              `json:"node"`
	Primary        bool                                `json:"primary"`
	RelocatingNode *string                             `json:"relocating_node,omitempty"`
	State          shardroutingstate.ShardRoutingState `json:"state"`
}

func (s *ShardRouting) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewShardRouting() *ShardRouting { _ = "STUB: not implemented"; return nil }
