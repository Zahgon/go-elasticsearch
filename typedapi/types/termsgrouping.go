package types

type TermsGrouping struct {
	Fields []string `json:"fields"`
}

func (s *TermsGrouping) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermsGrouping() *TermsGrouping { _ = "STUB: not implemented"; return nil }

type TermsGroupingVariant interface {
	TermsGroupingCaster() *TermsGrouping
}

func (s *TermsGrouping) TermsGroupingCaster() *TermsGrouping { _ = "STUB: not implemented"; return nil }
