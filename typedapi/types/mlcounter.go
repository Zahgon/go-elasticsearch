package types

type MlCounter struct {
	Count int64 `json:"count"`
}

func (s *MlCounter) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMlCounter() *MlCounter { _ = "STUB: not implemented"; return nil }
