package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _retentionPolicy struct {
	v *types.RetentionPolicy
}

func NewRetentionPolicy() *_retentionPolicy { _ = "STUB: not implemented"; return nil }

func (s *_retentionPolicy) Field(field string) *_retentionPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retentionPolicy) MaxAge(duration types.DurationVariant) *_retentionPolicy {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retentionPolicy) RetentionPolicyContainerCaster() *types.RetentionPolicyContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_retentionPolicy) RetentionPolicyCaster() *types.RetentionPolicy {
	_ = "STUB: not implemented"
	return nil
}
