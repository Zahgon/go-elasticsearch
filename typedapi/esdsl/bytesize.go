package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _byteSize struct {
	v types.ByteSize
}

func NewByteSize() *_byteSize { _ = "STUB: not implemented"; return nil }

func (u *_byteSize) Int64(int64 int64) *_byteSize { _ = "STUB: not implemented"; return nil }

func (u *_byteSize) String(string string) *_byteSize { _ = "STUB: not implemented"; return nil }

func (u *_byteSize) ByteSizeCaster() *types.ByteSize { _ = "STUB: not implemented"; return nil }
