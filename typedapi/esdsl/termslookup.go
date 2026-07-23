package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsLookup struct {
	v *types.TermsLookup
}

func NewTermsLookup() *_termsLookup { _ = "STUB: not implemented"; return nil }

func (s *_termsLookup) Id(id string) *_termsLookup { _ = "STUB: not implemented"; return nil }

func (s *_termsLookup) Index(indexname string) *_termsLookup { _ = "STUB: not implemented"; return nil }

func (s *_termsLookup) Path(field string) *_termsLookup { _ = "STUB: not implemented"; return nil }

func (s *_termsLookup) Routing(routing string) *_termsLookup { _ = "STUB: not implemented"; return nil }

func (s *_termsLookup) TermsLookupCaster() *types.TermsLookup {
	_ = "STUB: not implemented"
	return nil
}
