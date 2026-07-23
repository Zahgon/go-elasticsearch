package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/shardroutingstate"
)

type NodeShard struct {
	AllocationId          map[string]string                   `json:"allocation_id,omitempty"`
	Index                 string                              `json:"index"`
	Node                  *string                             `json:"node,omitempty"`
	Primary               bool                                `json:"primary"`
	RecoverySource        map[string]string                   `json:"recovery_source,omitempty"`
	RelocatingNode        *string                             `json:"relocating_node,omitempty"`
	RelocationFailureInfo *RelocationFailureInfo              `json:"relocation_failure_info,omitempty"`
	Shard                 int                                 `json:"shard"`
	State                 shardroutingstate.ShardRoutingState `json:"state"`
	UnassignedInfo        *UnassignedInformation              `json:"unassigned_info,omitempty"`
}

func (s *NodeShard) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewNodeShard() *NodeShard { _ = "STUB: not implemented"; return nil }
