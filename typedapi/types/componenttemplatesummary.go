package types

type ComponentTemplateSummary struct {
	Aliases           map[string]AliasDefinition `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptions         `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycle       `json:"lifecycle,omitempty"`
	Mappings          *TypeMapping               `json:"mappings,omitempty"`
	Meta_             Metadata                   `json:"_meta,omitempty"`
	Settings          map[string]IndexSettings   `json:"settings,omitempty"`
	Version           *int64                     `json:"version,omitempty"`
}

func (s *ComponentTemplateSummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewComponentTemplateSummary() *ComponentTemplateSummary { _ = "STUB: not implemented"; return nil }

type ComponentTemplateSummaryVariant interface {
	ComponentTemplateSummaryCaster() *ComponentTemplateSummary
}

func (s *ComponentTemplateSummary) ComponentTemplateSummaryCaster() *ComponentTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}
