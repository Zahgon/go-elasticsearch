package types

type SynonymsStats struct {
	Count      int `json:"count"`
	IndexCount int `json:"index_count"`
}

func (s *SynonymsStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSynonymsStats() *SynonymsStats { _ = "STUB: not implemented"; return nil }
