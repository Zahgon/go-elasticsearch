package types

type RankedDocument struct {
	Index          int     `json:"index"`
	RelevanceScore float32 `json:"relevance_score"`
	Text           *string `json:"text,omitempty"`
}

func (s *RankedDocument) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRankedDocument() *RankedDocument { _ = "STUB: not implemented"; return nil }
