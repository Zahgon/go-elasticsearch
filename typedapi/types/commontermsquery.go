package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/operator"
)

type CommonTermsQuery struct {
	Analyzer *string `json:"analyzer,omitempty"`

	Boost              *float32           `json:"boost,omitempty"`
	CutoffFrequency    *Float64           `json:"cutoff_frequency,omitempty"`
	HighFreqOperator   *operator.Operator `json:"high_freq_operator,omitempty"`
	LowFreqOperator    *operator.Operator `json:"low_freq_operator,omitempty"`
	MinimumShouldMatch MinimumShouldMatch `json:"minimum_should_match,omitempty"`
	Query              string             `json:"query"`
	QueryName_         *string            `json:"_name,omitempty"`
}

func (s *CommonTermsQuery) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCommonTermsQuery() *CommonTermsQuery { _ = "STUB: not implemented"; return nil }

type CommonTermsQueryVariant interface {
	CommonTermsQueryCaster() *CommonTermsQuery
}

func (s *CommonTermsQuery) CommonTermsQueryCaster() *CommonTermsQuery {
	_ = "STUB: not implemented"
	return nil
}
