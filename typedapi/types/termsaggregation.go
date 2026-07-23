package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationcollectmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/termsaggregationexecutionhint"
)

type TermsAggregation struct {
	CollectMode *termsaggregationcollectmode.TermsAggregationCollectMode `json:"collect_mode,omitempty"`

	Exclude []string `json:"exclude,omitempty"`

	ExecutionHint *termsaggregationexecutionhint.TermsAggregationExecutionHint `json:"execution_hint,omitempty"`

	Field  *string `json:"field,omitempty"`
	Format *string `json:"format,omitempty"`

	Include TermsInclude `json:"include,omitempty"`

	MinDocCount *int `json:"min_doc_count,omitempty"`

	Missing       Missing                    `json:"missing,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`

	Order  AggregateOrder `json:"order,omitempty"`
	Script *Script        `json:"script,omitempty"`

	ShardMinDocCount *int64 `json:"shard_min_doc_count,omitempty"`

	ShardSize *int `json:"shard_size,omitempty"`

	ShowTermDocCountError *bool `json:"show_term_doc_count_error,omitempty"`

	Size *int `json:"size,omitempty"`

	ValueType *string `json:"value_type,omitempty"`
}

func (s *TermsAggregation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermsAggregation() *TermsAggregation { _ = "STUB: not implemented"; return nil }

type TermsAggregationVariant interface {
	TermsAggregationCaster() *TermsAggregation
}

func (s *TermsAggregation) TermsAggregationCaster() *TermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
