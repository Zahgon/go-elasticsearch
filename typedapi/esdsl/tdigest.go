package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tdigestexecutionhint"
)

type _tDigest struct {
	v *types.TDigest
}

func NewTDigest() *_tDigest { _ = "STUB: not implemented"; return nil }

func (s *_tDigest) Compression(compression int) *_tDigest { _ = "STUB: not implemented"; return nil }

func (s *_tDigest) ExecutionHint(executionhint tdigestexecutionhint.TDigestExecutionHint) *_tDigest {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tDigest) TDigestCaster() *types.TDigest { _ = "STUB: not implemented"; return nil }
