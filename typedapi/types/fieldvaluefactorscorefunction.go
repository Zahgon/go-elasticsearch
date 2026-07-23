package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldvaluefactormodifier"
)

type FieldValueFactorScoreFunction struct {
	Factor *Float64 `json:"factor,omitempty"`

	Field string `json:"field"`

	Missing *Float64 `json:"missing,omitempty"`

	Modifier *fieldvaluefactormodifier.FieldValueFactorModifier `json:"modifier,omitempty"`
}

func (s *FieldValueFactorScoreFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFieldValueFactorScoreFunction() *FieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

type FieldValueFactorScoreFunctionVariant interface {
	FieldValueFactorScoreFunctionCaster() *FieldValueFactorScoreFunction
}

func (s *FieldValueFactorScoreFunction) FieldValueFactorScoreFunctionCaster() *FieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
