package allocationexplain

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/decision"
)

type Response struct {
	AllocateExplanation          *string                           `json:"allocate_explanation,omitempty"`
	AllocationDelay              types.Duration                    `json:"allocation_delay,omitempty"`
	AllocationDelayInMillis      *int64                            `json:"allocation_delay_in_millis,omitempty"`
	CanAllocate                  *decision.Decision                `json:"can_allocate,omitempty"`
	CanMoveToOtherNode           *decision.Decision                `json:"can_move_to_other_node,omitempty"`
	CanRebalanceCluster          *decision.Decision                `json:"can_rebalance_cluster,omitempty"`
	CanRebalanceClusterDecisions []types.AllocationDecision        `json:"can_rebalance_cluster_decisions,omitempty"`
	CanRebalanceToOtherNode      *decision.Decision                `json:"can_rebalance_to_other_node,omitempty"`
	CanRemainDecisions           []types.AllocationDecision        `json:"can_remain_decisions,omitempty"`
	CanRemainOnCurrentNode       *decision.Decision                `json:"can_remain_on_current_node,omitempty"`
	ClusterInfo                  *types.ClusterInfo                `json:"cluster_info,omitempty"`
	ConfiguredDelay              types.Duration                    `json:"configured_delay,omitempty"`
	ConfiguredDelayInMillis      *int64                            `json:"configured_delay_in_millis,omitempty"`
	CurrentNode                  *types.CurrentNode                `json:"current_node,omitempty"`
	CurrentState                 string                            `json:"current_state"`
	Index                        string                            `json:"index"`
	MoveExplanation              *string                           `json:"move_explanation,omitempty"`
	NodeAllocationDecisions      []types.NodeAllocationExplanation `json:"node_allocation_decisions,omitempty"`
	Note                         *string                           `json:"note,omitempty"`
	Primary                      bool                              `json:"primary"`
	RebalanceExplanation         *string                           `json:"rebalance_explanation,omitempty"`
	RemainingDelay               types.Duration                    `json:"remaining_delay,omitempty"`
	RemainingDelayInMillis       *int64                            `json:"remaining_delay_in_millis,omitempty"`
	Shard                        int                               `json:"shard"`
	UnassignedInfo               *types.UnassignedInformation      `json:"unassigned_info,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
