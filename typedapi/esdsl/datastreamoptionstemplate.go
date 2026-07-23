package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _dataStreamOptionsTemplate struct {
	v *types.DataStreamOptionsTemplate
}

func NewDataStreamOptionsTemplate() *_dataStreamOptionsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamOptionsTemplate) FailureStore(failurestore types.DataStreamFailureStoreTemplateVariant) *_dataStreamOptionsTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (s *_dataStreamOptionsTemplate) DataStreamOptionsTemplateCaster() *types.DataStreamOptionsTemplate {
	_ = "STUB: not implemented"
	return nil
}
