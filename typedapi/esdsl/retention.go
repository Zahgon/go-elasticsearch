package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _retention struct {
	v *types.Retention
}

func NewRetention(maxcount int, mincount int) *_retention { _ = "STUB: not implemented"; return nil }

func (s *_retention) ExpireAfter(duration types.DurationVariant) *_retention {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retention) MaxCount(maxcount int) *_retention { _ = "STUB: not implemented"; return nil }

func (s *_retention) MinCount(mincount int) *_retention { _ = "STUB: not implemented"; return nil }

func (s *_retention) RetentionCaster() *types.Retention { _ = "STUB: not implemented"; return nil }
