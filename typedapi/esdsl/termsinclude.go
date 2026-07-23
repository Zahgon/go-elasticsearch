package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _termsInclude struct {
	v types.TermsInclude
}

func NewTermsInclude() *_termsInclude { _ = "STUB: not implemented"; return nil }

func (u *_termsInclude) String(string string) *_termsInclude { _ = "STUB: not implemented"; return nil }

func (u *_termsInclude) Strings(strings ...string) *_termsInclude {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsInclude) TermsPartition(termspartition types.TermsPartitionVariant) *_termsInclude {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsPartition) TermsIncludeCaster() *types.TermsInclude {
	_ = "STUB: not implemented"
	return nil
}

func (u *_termsInclude) TermsIncludeCaster() *types.TermsInclude {
	_ = "STUB: not implemented"
	return nil
}
