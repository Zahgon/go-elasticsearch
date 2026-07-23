package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditionop"
)

type ArrayCompareCondition struct {
	ArrayCompareCondition map[conditionop.ConditionOp]ArrayCompareOpParams `json:"-"`
	Path                  string                                           `json:"path"`
}

func (s *ArrayCompareCondition) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func (s ArrayCompareCondition) MarshalJSON() ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewArrayCompareCondition() *ArrayCompareCondition { _ = "STUB: not implemented"; return nil }

type ArrayCompareConditionVariant interface {
	ArrayCompareConditionCaster() *ArrayCompareCondition
}

func (s *ArrayCompareCondition) ArrayCompareConditionCaster() *ArrayCompareCondition {
	_ = "STUB: not implemented"
	return nil
}
