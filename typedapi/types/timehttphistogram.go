package types

type TimeHttpHistogram struct {
	Count    int64  `json:"count"`
	GeMillis *int64 `json:"ge_millis,omitempty"`
	LtMillis *int64 `json:"lt_millis,omitempty"`
}

func (s *TimeHttpHistogram) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTimeHttpHistogram() *TimeHttpHistogram { _ = "STUB: not implemented"; return nil }
