package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/fieldvaluefactormodifier"
)

type _fieldValueFactorScoreFunction struct {
	v *types.FieldValueFactorScoreFunction
}

func NewFieldValueFactorScoreFunction() *_fieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) Factor(factor types.Float64) *_fieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) Field(field string) *_fieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) Missing(missing types.Float64) *_fieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) Modifier(modifier fieldvaluefactormodifier.FieldValueFactorModifier) *_fieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) FunctionScoreCaster() *types.FunctionScore {
	_ = "STUB: not implemented"
	return nil
}

func (s *_fieldValueFactorScoreFunction) FieldValueFactorScoreFunctionCaster() *types.FieldValueFactorScoreFunction {
	_ = "STUB: not implemented"
	return nil
}
