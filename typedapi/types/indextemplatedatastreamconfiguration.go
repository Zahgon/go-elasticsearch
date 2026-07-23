package types

type IndexTemplateDataStreamConfiguration struct {
	AllowCustomRouting *bool `json:"allow_custom_routing,omitempty"`

	Hidden *bool `json:"hidden,omitempty"`
}

func (s *IndexTemplateDataStreamConfiguration) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexTemplateDataStreamConfiguration() *IndexTemplateDataStreamConfiguration {
	_ = "STUB: not implemented"
	return nil
}

type IndexTemplateDataStreamConfigurationVariant interface {
	IndexTemplateDataStreamConfigurationCaster() *IndexTemplateDataStreamConfiguration
}

func (s *IndexTemplateDataStreamConfiguration) IndexTemplateDataStreamConfigurationCaster() *IndexTemplateDataStreamConfiguration {
	_ = "STUB: not implemented"
	return nil
}
