package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/missingorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/sortorder"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/valuetype"
)

type CompositeTermsAggregation struct {
	Field         *string                    `json:"field,omitempty"`
	MissingBucket *bool                      `json:"missing_bucket,omitempty"`
	MissingOrder  *missingorder.MissingOrder `json:"missing_order,omitempty"`
	Order         *sortorder.SortOrder       `json:"order,omitempty"`

	Script    *Script              `json:"script,omitempty"`
	ValueType *valuetype.ValueType `json:"value_type,omitempty"`
}

func (s *CompositeTermsAggregation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompositeTermsAggregation() *CompositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}

type CompositeTermsAggregationVariant interface {
	CompositeTermsAggregationCaster() *CompositeTermsAggregation
}

func (s *CompositeTermsAggregation) CompositeTermsAggregationCaster() *CompositeTermsAggregation {
	_ = "STUB: not implemented"
	return nil
}
