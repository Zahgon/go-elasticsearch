package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
)

type MultiTermsAggregation struct {
	CollectMode *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`

	MinDocCount *int64 `json:"min_doc_count,omitempty"`

	Order AggregateOrder `json:"order,omitempty"`

	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	ShowTermDocCountError *bool `json:"show_term_doc_count_error,omitempty"`

	Size *int `json:"size,omitempty"`

	Terms []MultiTermLookup `json:"terms"`
}

func (s *MultiTermsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMultiTermsAggregation() *MultiTermsAggregation { _ = "STUB: not implemented"; return nil }

type MultiTermsAggregationVariant interface {
	MultiTermsAggregationCaster() *MultiTermsAggregation
}

func (s *MultiTermsAggregation) MultiTermsAggregationCaster() *MultiTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
