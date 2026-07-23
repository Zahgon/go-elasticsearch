package types

type WatchStatus struct {
	Actions          WatcherStatusActions `json:"actions"`
	ExecutionState   *string              `json:"execution_state,omitempty"`
	LastChecked      DateTime             `json:"last_checked,omitempty"`
	LastMetCondition DateTime             `json:"last_met_condition,omitempty"`
	State            ActivationState      `json:"state"`
	Version          int64                `json:"version"`
}

func (s *WatchStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewWatchStatus() *WatchStatus { _ = "STUB: not implemented"; return nil }

type WatchStatusVariant interface {
	WatchStatusCaster() *WatchStatus
}

func (s *WatchStatus) WatchStatusCaster() *WatchStatus { _ = "STUB: not implemented"; return nil }
