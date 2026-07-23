package types

type CCSUsageTimeValue struct {
	Avg int64 `json:"avg"`

	Max int64 `json:"max"`

	P90 int64 `json:"p90"`
}

func (s *CCSUsageTimeValue) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCCSUsageTimeValue() *CCSUsageTimeValue { _ = "STUB: not implemented"; return nil }
