package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _softDeletes struct {
	v *types.SoftDeletes
}

func NewSoftDeletes() *_softDeletes { _ = "STUB: not implemented"; return nil }

func (s *_softDeletes) Enabled(enabled bool) *_softDeletes { _ = "STUB: not implemented"; return nil }

func (s *_softDeletes) RetentionLease(retentionlease types.RetentionLeaseVariant) *_softDeletes {
	_ = "STUB: not implemented"
	return nil
}

func (s *_softDeletes) SoftDeletesCaster() *types.SoftDeletes {
	_ = "STUB: not implemented"
	return nil
}
