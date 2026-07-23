package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/allocationexplaindecision"
)

type AllocationDecision struct {
	Decider     string                                              `json:"decider"`
	Decision    allocationexplaindecision.AllocationExplainDecision `json:"decision"`
	Explanation string                                              `json:"explanation"`
}

func (s *AllocationDecision) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAllocationDecision() *AllocationDecision { _ = "STUB: not implemented"; return nil }
