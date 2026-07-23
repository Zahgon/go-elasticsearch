package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsExclude struct {
	v types.TermsExclude
}

func NewTermsExclude() *_termsExclude { _ = "STUB: not implemented"; return nil }

func (u *_termsExclude) Strings(strings ...string) *_termsExclude {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsExclude) TermsExcludeCaster() *types.TermsExclude {
	_ = "STUB: not implemented"
	return nil
}
