package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _bucketsPath struct {
	v types.BucketsPath
}

func NewBucketsPath() *_bucketsPath { _ = "STUB: not implemented"; return nil }

func (u *_bucketsPath) String(string string) *_bucketsPath { _ = "STUB: not implemented"; return nil }

func (u *_bucketsPath) Strings(strings ...string) *_bucketsPath {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bucketsPath) Map(value map[string]string) *_bucketsPath {
	_ = "STUB: not implemented"
	return nil
}

func (u *_bucketsPath) BucketsPathCaster() *types.BucketsPath {
	_ = "STUB: not implemented"
	return nil
}
