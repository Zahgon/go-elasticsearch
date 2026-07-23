package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/quantifier"
)

type ArrayCompareOpParams struct {
	Quantifier quantifier.Quantifier `json:"quantifier"`
	Value      FieldValue            `json:"value"`
}

func (s *ArrayCompareOpParams) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewArrayCompareOpParams() *ArrayCompareOpParams { _ = "STUB: not implemented"; return nil }

type ArrayCompareOpParamsVariant interface {
	ArrayCompareOpParamsCaster() *ArrayCompareOpParams
}

func (s *ArrayCompareOpParams) ArrayCompareOpParamsCaster() *ArrayCompareOpParams {
	_ = "STUB: not implemented"
	return nil
}
