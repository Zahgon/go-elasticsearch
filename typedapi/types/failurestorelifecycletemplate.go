package types

type FailureStoreLifecycleTemplate struct {
	DataRetention Duration `json:"data_retention,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

func (s *FailureStoreLifecycleTemplate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFailureStoreLifecycleTemplate() *FailureStoreLifecycleTemplate {
	_ = "STUB: not implemented"
	return nil
}

type FailureStoreLifecycleTemplateVariant interface {
	FailureStoreLifecycleTemplateCaster() *FailureStoreLifecycleTemplate
}

func (s *FailureStoreLifecycleTemplate) FailureStoreLifecycleTemplateCaster() *FailureStoreLifecycleTemplate {
	_ = "STUB: not implemented"
	return nil
}
