package types

type TimeSync struct {
	Delay Duration `json:"delay,omitempty"`

	Field string `json:"field"`
}

func (s *TimeSync) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTimeSync() *TimeSync { _ = "STUB: not implemented"; return nil }

type TimeSyncVariant interface {
	TimeSyncCaster() *TimeSync
}

func (s *TimeSync) TimeSyncCaster() *TimeSync { _ = "STUB: not implemented"; return nil }
