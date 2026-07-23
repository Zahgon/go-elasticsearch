package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamFailureStoreTemplate struct {
	v *types.DataStreamFailureStoreTemplate
}

func NewDataStreamFailureStoreTemplate() *_dataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamFailureStoreTemplate) Enabled(enabled bool) *_dataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamFailureStoreTemplate) Lifecycle(lifecycle types.FailureStoreLifecycleTemplateVariant) *_dataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamFailureStoreTemplate) DataStreamFailureStoreTemplateCaster() *types.DataStreamFailureStoreTemplate {
	_ = "STUB: not implemented"
	return nil
}
