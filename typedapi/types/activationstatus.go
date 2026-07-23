package types

type ActivationStatus struct {
	Actions WatcherStatusActions `json:"actions"`
	State   ActivationState      `json:"state"`
	Version int64                `json:"version"`
}

func (s *ActivationStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewActivationStatus() *ActivationStatus { _ = "STUB: not implemented"; return nil }
