package types

type IndexTemplateMapping struct {
	Aliases           map[string]Alias           `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptionsTemplate `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycle       `json:"lifecycle,omitempty"`

	Mappings *TypeMapping `json:"mappings,omitempty"`

	Settings *IndexSettings `json:"settings,omitempty"`
}

func (s *IndexTemplateMapping) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexTemplateMapping() *IndexTemplateMapping { _ = "STUB: not implemented"; return nil }

type IndexTemplateMappingVariant interface {
	IndexTemplateMappingCaster() *IndexTemplateMapping
}

func (s *IndexTemplateMapping) IndexTemplateMappingCaster() *IndexTemplateMapping {
	_ = "STUB: not implemented"
	return nil
}
