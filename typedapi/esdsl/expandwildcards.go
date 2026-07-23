package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _expandWildcards struct {
	v types.ExpandWildcards
}

func NewExpandWildcards() *_expandWildcards { _ = "STUB: not implemented"; return nil }

func (u *_expandWildcards) ExpandWildcardsCaster() *types.ExpandWildcards {
	_ = "STUB: not implemented"
	return nil
}
