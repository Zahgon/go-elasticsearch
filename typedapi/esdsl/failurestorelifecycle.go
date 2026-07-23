package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _failureStoreLifecycle struct {
	v *types.FailureStoreLifecycle
}

func NewFailureStoreLifecycle() *_failureStoreLifecycle { _ = "STUB: not implemented"; return nil }

func (s *_failureStoreLifecycle) DataRetention(duration types.DurationVariant) *_failureStoreLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failureStoreLifecycle) Enabled(enabled bool) *_failureStoreLifecycle {
	_ = "STUB: not implemented"
	return nil
}

func (s *_failureStoreLifecycle) FailureStoreLifecycleCaster() *types.FailureStoreLifecycle {
	_ = "STUB: not implemented"
	return nil
}
