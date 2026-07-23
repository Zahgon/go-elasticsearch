package putwatch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Actions map[string]types.WatcherAction `json:"actions,omitempty"`

	Condition *types.WatcherCondition `json:"condition,omitempty"`

	Input *types.WatcherInput `json:"input,omitempty"`

	Metadata types.Metadata `json:"metadata,omitempty"`

	ThrottlePeriod types.Duration `json:"throttle_period,omitempty"`

	ThrottlePeriodInMillis *int64 `json:"throttle_period_in_millis,omitempty"`

	Transform *types.TransformContainer `json:"transform,omitempty"`

	Trigger *types.TriggerContainer `json:"trigger,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
