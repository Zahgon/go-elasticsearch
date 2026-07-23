package types

type FailureStore struct {
	Enabled         bool              `json:"enabled"`
	Indices         []DataStreamIndex `json:"indices"`
	RolloverOnWrite bool              `json:"rollover_on_write"`
}

func (s *FailureStore) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewFailureStore() *FailureStore { _ = "STUB: not implemented"; return nil }
