package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _missing struct {
	v types.Missing
}

func NewMissing() *_missing { _ = "STUB: not implemented"; return nil }

func (u *_missing) String(string string) *_missing { _ = "STUB: not implemented"; return nil }

func (u *_missing) Int(int int) *_missing { _ = "STUB: not implemented"; return nil }

func (u *_missing) Float64(float64 types.Float64) *_missing { _ = "STUB: not implemented"; return nil }

func (u *_missing) Bool(bool bool) *_missing { _ = "STUB: not implemented"; return nil }

func (u *_missing) MissingCaster() *types.Missing { _ = "STUB: not implemented"; return nil }
