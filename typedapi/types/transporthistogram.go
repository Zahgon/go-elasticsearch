package types

type TransportHistogram struct {
	Count *int64 `json:"count,omitempty"`

	GeMillis *int64 `json:"ge_millis,omitempty"`

	LtMillis *int64 `json:"lt_millis,omitempty"`
}

func (s *TransportHistogram) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTransportHistogram() *TransportHistogram { _ = "STUB: not implemented"; return nil }
