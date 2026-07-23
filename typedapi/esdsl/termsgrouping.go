package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsGrouping struct {
	v *types.TermsGrouping
}

func NewTermsGrouping() *_termsGrouping { _ = "STUB: not implemented"; return nil }

func (s *_termsGrouping) Fields(fields ...string) *_termsGrouping {
	_ = "STUB: not implemented"
	return nil
}

func (s *_termsGrouping) TermsGroupingCaster() *types.TermsGrouping {
	_ = "STUB: not implemented"
	return nil
}
