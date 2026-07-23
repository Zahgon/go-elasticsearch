package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/acknowledgementoptions"
)

type AcknowledgeState struct {
	State     acknowledgementoptions.AcknowledgementOptions `json:"state"`
	Timestamp DateTime                                      `json:"timestamp"`
}

func (s *AcknowledgeState) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAcknowledgeState() *AcknowledgeState { _ = "STUB: not implemented"; return nil }

type AcknowledgeStateVariant interface {
	AcknowledgeStateCaster() *AcknowledgeState
}

func (s *AcknowledgeState) AcknowledgeStateCaster() *AcknowledgeState {
	_ = "STUB: not implemented"
	return nil
}
