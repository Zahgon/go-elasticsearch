package types

type IndexTemplate struct {
	AllowAutoCreate *bool `json:"allow_auto_create,omitempty"`

	ComposedOf []string `json:"composed_of"`

	CreatedDate DateTime `json:"created_date,omitempty"`

	CreatedDateMillis *int64 `json:"created_date_millis,omitempty"`

	DataStream *IndexTemplateDataStreamConfiguration `json:"data_stream,omitempty"`

	Deprecated *bool `json:"deprecated,omitempty"`

	IgnoreMissingComponentTemplates []string `json:"ignore_missing_component_templates,omitempty"`

	IndexPatterns []string `json:"index_patterns"`

	Meta_ Metadata `json:"_meta,omitempty"`

	ModifiedDate DateTime `json:"modified_date,omitempty"`

	ModifiedDateMillis *int64 `json:"modified_date_millis,omitempty"`

	Priority *int64 `json:"priority,omitempty"`

	Template *IndexTemplateSummary `json:"template,omitempty"`

	Version *int64 `json:"version,omitempty"`
}

func (s *IndexTemplate) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewIndexTemplate() *IndexTemplate { _ = "STUB: not implemented"; return nil }

type IndexTemplateVariant interface {
	IndexTemplateCaster() *IndexTemplate
}

func (s *IndexTemplate) IndexTemplateCaster() *IndexTemplate { _ = "STUB: not implemented"; return nil }
