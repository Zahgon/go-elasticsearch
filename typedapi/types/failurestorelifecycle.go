package types

type FailureStoreLifecycle struct {
	DataRetention Duration `json:"data_retention,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

func (s *FailureStoreLifecycle) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFailureStoreLifecycle() *FailureStoreLifecycle { _ = "STUB: not implemented"; return nil }

type FailureStoreLifecycleVariant interface {
	FailureStoreLifecycleCaster() *FailureStoreLifecycle
}

func (s *FailureStoreLifecycle) FailureStoreLifecycleCaster() *FailureStoreLifecycle {
	_ = "STUB: not implemented"
	return nil
}
