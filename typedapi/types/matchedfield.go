package types

type MatchedField struct {
	Length int    `json:"length"`
	Match  string `json:"match"`
	Offset int    `json:"offset"`
}

func (s *MatchedField) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMatchedField() *MatchedField { _ = "STUB: not implemented"; return nil }
