package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningeffort"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/reasoningsummary"
)

type Reasoning struct {
	Effort *reasoningeffort.ReasoningEffort `json:"effort,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`

	Exclude *bool `json:"exclude,omitempty"`

	Summary *reasoningsummary.ReasoningSummary `json:"summary,omitempty"`
}

func (s *Reasoning) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewReasoning() *Reasoning { _ = "STUB: not implemented"; return nil }

type ReasoningVariant interface {
	ReasoningCaster() *Reasoning
}

func (s *Reasoning) ReasoningCaster() *Reasoning { _ = "STUB: not implemented"; return nil }
