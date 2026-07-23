package types

type SizeHttpHistogram struct {
	Count   int64  `json:"count"`
	GeBytes *int64 `json:"ge_bytes,omitempty"`
	LtBytes *int64 `json:"lt_bytes,omitempty"`
}

func (s *SizeHttpHistogram) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSizeHttpHistogram() *SizeHttpHistogram { _ = "STUB: not implemented"; return nil }
