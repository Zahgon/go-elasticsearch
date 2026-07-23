package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/decision"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noderole"
)

type NodeAllocationExplanation struct {
	Deciders         []AllocationDecision `json:"deciders,omitempty"`
	NodeAttributes   map[string]string    `json:"node_attributes"`
	NodeDecision     decision.Decision    `json:"node_decision"`
	NodeId           string               `json:"node_id"`
	NodeName         string               `json:"node_name"`
	Roles            []noderole.NodeRole  `json:"roles"`
	Store            *AllocationStore     `json:"store,omitempty"`
	TransportAddress string               `json:"transport_address"`
	WeightRanking    *int                 `json:"weight_ranking,omitempty"`
}

func (s *NodeAllocationExplanation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewNodeAllocationExplanation() *NodeAllocationExplanation {
	_ = "STUB: not implemented"
	return nil
}
