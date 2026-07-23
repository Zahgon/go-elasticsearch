package types

type StreamStatus struct {
	Enabled bool `json:"enabled"`
}

func (s *StreamStatus) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStreamStatus() *StreamStatus { _ = "STUB: not implemented"; return nil }
