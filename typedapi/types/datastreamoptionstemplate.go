package types

type DataStreamOptionsTemplate struct {
	FailureStore *DataStreamFailureStoreTemplate `json:"failure_store,omitempty"`
}

func (s *DataStreamOptionsTemplate) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataStreamOptionsTemplate() *DataStreamOptionsTemplate {
	_ = "STUB: not implemented"
	return nil
}

type DataStreamOptionsTemplateVariant interface {
	DataStreamOptionsTemplateCaster() *DataStreamOptionsTemplate
}

func (s *DataStreamOptionsTemplate) DataStreamOptionsTemplateCaster() *DataStreamOptionsTemplate {
	_ = "STUB: not implemented"
	return nil
}
