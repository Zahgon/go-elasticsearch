package types

type SynonymsSetItem struct {
	Count int `json:"count"`

	SynonymsSet string `json:"synonyms_set"`
}

func (s *SynonymsSetItem) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSynonymsSetItem() *SynonymsSetItem { _ = "STUB: not implemented"; return nil }
