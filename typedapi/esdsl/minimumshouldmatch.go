package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _minimumShouldMatch struct {
	v types.MinimumShouldMatch
}

func NewMinimumShouldMatch() *_minimumShouldMatch { _ = "STUB: not implemented"; return nil }

func (u *_minimumShouldMatch) Int(int int) *_minimumShouldMatch {
	_ = "STUB: not implemented"
	return nil
}

func (u *_minimumShouldMatch) String(string string) *_minimumShouldMatch {
	_ = "STUB: not implemented"
	return nil
}

func (u *_minimumShouldMatch) MinimumShouldMatchCaster() *types.MinimumShouldMatch {
	_ = "STUB: not implemented"
	return nil
}
