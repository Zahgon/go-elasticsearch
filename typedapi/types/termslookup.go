package types

type TermsLookup struct {
	Id      string  `json:"id"`
	Index   string  `json:"index"`
	Path    string  `json:"path"`
	Routing *string `json:"routing,omitempty"`
}

func (s *TermsLookup) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewTermsLookup() *TermsLookup { _ = "STUB: not implemented"; return nil }

type TermsLookupVariant interface {
	TermsLookupCaster() *TermsLookup
}

func (s *TermsLookup) TermsLookupCaster() *TermsLookup { _ = "STUB: not implemented"; return nil }

func (s *TermsLookup) TermsQueryFieldCaster() *TermsQueryField {
	_ = "STUB: not implemented"
	return nil
}
