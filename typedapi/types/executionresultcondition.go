package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/conditiontype"
)

type ExecutionResultCondition struct {
	Met    bool                                    `json:"met"`
	Status actionstatusoptions.ActionStatusOptions `json:"status"`
	Type   conditiontype.ConditionType             `json:"type"`
}

func (s *ExecutionResultCondition) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewExecutionResultCondition() *ExecutionResultCondition { _ = "STUB: not implemented"; return nil }
