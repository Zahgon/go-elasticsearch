package types

type SynonymRule struct {
	Id *string `json:"id,omitempty"`

	Synonyms string `json:"synonyms"`
}

func (s *SynonymRule) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSynonymRule() *SynonymRule { _ = "STUB: not implemented"; return nil }

type SynonymRuleVariant interface {
	SynonymRuleCaster() *SynonymRule
}

func (s *SynonymRule) SynonymRuleCaster() *SynonymRule { _ = "STUB: not implemented"; return nil }
