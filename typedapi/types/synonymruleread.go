package types

type SynonymRuleRead struct {
	Id string `json:"id"`

	Synonyms string `json:"synonyms"`
}

func (s *SynonymRuleRead) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSynonymRuleRead() *SynonymRuleRead { _ = "STUB: not implemented"; return nil }
