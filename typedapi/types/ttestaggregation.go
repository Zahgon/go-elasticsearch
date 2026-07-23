package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/ttesttype"
)

type TTestAggregation struct {
	A *TestPopulation `json:"a,omitempty"`

	B *TestPopulation `json:"b,omitempty"`

	Type *ttesttype.TTestType `json:"type,omitempty"`
}

func NewTTestAggregation() *TTestAggregation { _ = "STUB: not implemented"; return nil }

type TTestAggregationVariant interface {
	TTestAggregationCaster() *TTestAggregation
}

func (s *TTestAggregation) TTestAggregationCaster() *TTestAggregation {
	_ = "STUB: not implemented"
	return nil
}
