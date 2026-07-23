package types

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionstatusoptions"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/inputtype"
)

type ExecutionResultInput struct {
	Payload map[string]json.RawMessage              `json:"payload"`
	Status  actionstatusoptions.ActionStatusOptions `json:"status"`
	Type    inputtype.InputType                     `json:"type"`
}

func NewExecutionResultInput() *ExecutionResultInput { _ = "STUB: not implemented"; return nil }
