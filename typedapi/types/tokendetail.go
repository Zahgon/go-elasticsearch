package types

type TokenDetail struct {
	Name   string                `json:"name"`
	Tokens []ExplainAnalyzeToken `json:"tokens"`
}

func (s *TokenDetail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTokenDetail() *TokenDetail { _ = "STUB: not implemented"; return nil }
