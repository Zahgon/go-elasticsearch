package types

type Watch struct {
	Actions                map[string]WatcherAction `json:"actions"`
	Condition              WatcherCondition         `json:"condition"`
	Input                  WatcherInput             `json:"input"`
	Metadata               Metadata                 `json:"metadata,omitempty"`
	Status                 *WatchStatus             `json:"status,omitempty"`
	ThrottlePeriod         Duration                 `json:"throttle_period,omitempty"`
	ThrottlePeriodInMillis *int64                   `json:"throttle_period_in_millis,omitempty"`
	Transform              *TransformContainer      `json:"transform,omitempty"`
	Trigger                TriggerContainer         `json:"trigger"`
}

func (s *Watch) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatch() *Watch { _ = "STUB: not implemented"; return nil }

type WatchVariant interface {
	WatchCaster() *Watch
}

func (s *Watch) WatchCaster() *Watch { _ = "STUB: not implemented"; return nil }
