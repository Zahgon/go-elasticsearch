package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/enrichpolicyphase"
)

type ExecuteEnrichPolicyStatus struct {
	Phase enrichpolicyphase.EnrichPolicyPhase `json:"phase"`
	Step  *string                             `json:"step,omitempty"`
}

func (s *ExecuteEnrichPolicyStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExecuteEnrichPolicyStatus() *ExecuteEnrichPolicyStatus {
	_ = "STUB: not implemented"
	return nil
}
