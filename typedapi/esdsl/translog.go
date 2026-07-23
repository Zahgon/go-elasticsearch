package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/translogdurability"
)

type _translog struct {
	v *types.Translog
}

func NewTranslog() *_translog { _ = "STUB: not implemented"; return nil }

func (s *_translog) Durability(durability translogdurability.TranslogDurability) *_translog {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translog) FlushThresholdSize(bytesize types.ByteSizeVariant) *_translog {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translog) Retention(retention types.TranslogRetentionVariant) *_translog {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translog) SyncInterval(duration types.DurationVariant) *_translog {
	_ = "STUB: not implemented"
	return nil
}

func (s *_translog) TranslogCaster() *types.Translog { _ = "STUB: not implemented"; return nil }
