package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsQueryField struct {
	v types.TermsQueryField
}

func NewTermsQueryField() *_termsQueryField { _ = "STUB: not implemented"; return nil }

func (u *_termsQueryField) FieldValues(fieldvalues ...types.FieldValueVariant) *_termsQueryField {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsQueryField) TermsLookup(termslookup types.TermsLookupVariant) *_termsQueryField {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsLookup) TermsQueryFieldCaster() *types.TermsQueryField {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsQueryField) TermsQueryFieldCaster() *types.TermsQueryField {
	_ = "STUB: not implemented"
	return nil
}
