package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _translogRetention struct {
	v *types.TranslogRetention
}

func NewTranslogRetention() *_translogRetention { _ = "STUB: not implemented"; return nil }

func (s *_translogRetention) Age(duration types.DurationVariant) *_translogRetention {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translogRetention) Size(bytesize types.ByteSizeVariant) *_translogRetention {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translogRetention) TranslogRetentionCaster() *types.TranslogRetention {
	_ = "STUB: not implemented"
	return nil
}
