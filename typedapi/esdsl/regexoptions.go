package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _regexOptions struct {
	v *types.RegexOptions
}

func NewRegexOptions() *_regexOptions { _ = "STUB: not implemented"; return nil }

func (s *_regexOptions) Flags(flags string) *_regexOptions { _ = "STUB: not implemented"; return nil }

func (s *_regexOptions) MaxDeterminizedStates(maxdeterminizedstates int) *_regexOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_regexOptions) RegexOptionsCaster() *types.RegexOptions {
	_ = "STUB: not implemented"
	return nil
}
