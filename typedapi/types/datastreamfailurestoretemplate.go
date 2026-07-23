package types

type DataStreamFailureStoreTemplate struct {
	Enabled *bool `json:"enabled,omitempty"`

	Lifecycle *FailureStoreLifecycleTemplate `json:"lifecycle,omitempty"`
}

func (s *DataStreamFailureStoreTemplate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamFailureStoreTemplate() *DataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}

type DataStreamFailureStoreTemplateVariant interface {
	DataStreamFailureStoreTemplateCaster() *DataStreamFailureStoreTemplate
}

func (s *DataStreamFailureStoreTemplate) DataStreamFailureStoreTemplateCaster() *DataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}
