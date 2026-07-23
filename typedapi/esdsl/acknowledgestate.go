package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/acknowledgementoptions"
)

type _acknowledgeState struct {
	v *types.AcknowledgeState
}

func NewAcknowledgeState(state acknowledgementoptions.AcknowledgementOptions) *_acknowledgeState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_acknowledgeState) State(state acknowledgementoptions.AcknowledgementOptions) *_acknowledgeState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_acknowledgeState) Timestamp(datetime types.DateTimeVariant) *_acknowledgeState {
	_ = "STUB: not implemented"
	return nil
}

func (s *_acknowledgeState) AcknowledgeStateCaster() *types.AcknowledgeState {
	_ = "STUB: not implemented"
	return nil
}
