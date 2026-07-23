package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _stringifiedboolean struct {
	v types.Stringifiedboolean
}

func NewStringifiedboolean() *_stringifiedboolean { _ = "STUB: not implemented"; return nil }

func (u *_stringifiedboolean) Bool(bool bool) *_stringifiedboolean {
	_ = "STUB: not implemented"
	return nil
}

func (u *_stringifiedboolean) String(string string) *_stringifiedboolean {
	_ = "STUB: not implemented"
	return nil
}

func (u *_stringifiedboolean) StringifiedbooleanCaster() *types.Stringifiedboolean {
	_ = "STUB: not implemented"
	return nil
}
