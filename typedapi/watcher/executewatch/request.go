package executewatch

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/actionexecutionmode"
)

type Request struct {
	ActionModes map[string]actionexecutionmode.ActionExecutionMode `json:"action_modes,omitempty"`

	AlternativeInput map[string]json.RawMessage `json:"alternative_input,omitempty"`

	IgnoreCondition *bool `json:"ignore_condition,omitempty"`

	RecordExecution  *bool                   `json:"record_execution,omitempty"`
	SimulatedActions *types.SimulatedActions `json:"simulated_actions,omitempty"`

	TriggerData *types.ScheduleTriggerEvent `json:"trigger_data,omitempty"`

	Watch *types.Watch `json:"watch,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
