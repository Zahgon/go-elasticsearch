package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _ids struct {
	v types.Ids
}

func NewIds() *_ids { _ = "STUB: not implemented"; return nil }

func (u *_ids) IdsCaster() *types.Ids { _ = "STUB: not implemented"; return nil }
