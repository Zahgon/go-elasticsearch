package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionop"
)

type _arrayCompareCondition struct {
	k string
	v *types.ArrayCompareCondition
}

func NewArrayCompareCondition(key string) *_arrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareCondition) ArrayCompareCondition(arraycomparecondition map[conditionop.ConditionOp]types.ArrayCompareOpParams) *_arrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareCondition) AddArrayCompareCondition(key conditionop.ConditionOp, value types.ArrayCompareOpParamsVariant) *_arrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareCondition) Path(path string) *_arrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareCondition) WatcherConditionCaster() *types.WatcherCondition {
	_ = "STUB: not implemented"
	return nil
}

func NewSingleArrayCompareCondition() *_arrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}

func (s *_arrayCompareCondition) ArrayCompareConditionCaster() *types.ArrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}
