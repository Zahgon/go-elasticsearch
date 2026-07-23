package types

type AddIndicesBlockStatus struct {
	Blocked bool   `json:"blocked"`
	Name    string `json:"name"`
}

func (s *AddIndicesBlockStatus) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewAddIndicesBlockStatus() *AddIndicesBlockStatus { _ = "STUB: not implemented"; return nil }
