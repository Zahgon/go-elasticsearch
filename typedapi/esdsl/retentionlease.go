package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _retentionLease struct {
	v *types.RetentionLease
}

func NewRetentionLease() *_retentionLease { _ = "STUB: not implemented"; return nil }

func (s *_retentionLease) Period(duration types.DurationVariant) *_retentionLease {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retentionLease) RetentionLeaseCaster() *types.RetentionLease {
	_ = "STUB: not implemented"
	return nil
}
