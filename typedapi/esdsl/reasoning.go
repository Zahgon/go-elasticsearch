package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningeffort"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningsummary"
)

type _reasoning struct {
	v *types.Reasoning
}

func NewReasoning() *_reasoning { _ = "STUB: not implemented"; return nil }

func (s *_reasoning) Effort(effort reasoningeffort.ReasoningEffort) *_reasoning {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reasoning) Enabled(enabled bool) *_reasoning { _ = "STUB: not implemented"; return nil }

func (s *_reasoning) Exclude(exclude bool) *_reasoning { _ = "STUB: not implemented"; return nil }

func (s *_reasoning) Summary(summary reasoningsummary.ReasoningSummary) *_reasoning {
	_ = "STUB: not implemented"
	return nil
}

func (s *_reasoning) ReasoningCaster() *types.Reasoning { _ = "STUB: not implemented"; return nil }
