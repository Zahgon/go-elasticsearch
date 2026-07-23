package types

type CountRecord struct {
	Count *string `json:"count,omitempty"`

	Epoch StringifiedEpochTimeUnitSeconds `json:"epoch,omitempty"`

	Timestamp *string `json:"timestamp,omitempty"`
}

func (s *CountRecord) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCountRecord() *CountRecord { _ = "STUB: not implemented"; return nil }
