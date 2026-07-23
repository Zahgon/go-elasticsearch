package types

type IndexTemplateSummary struct {
	Aliases           map[string]Alias     `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptions   `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycle `json:"lifecycle,omitempty"`

	Mappings *TypeMapping `json:"mappings,omitempty"`

	Settings *IndexSettings `json:"settings,omitempty"`
}

func NewIndexTemplateSummary() *IndexTemplateSummary { _ = "STUB: not implemented"; return nil }

type IndexTemplateSummaryVariant interface {
	IndexTemplateSummaryCaster() *IndexTemplateSummary
}

func (s *IndexTemplateSummary) IndexTemplateSummaryCaster() *IndexTemplateSummary {
	_ = "STUB: not implemented"
	return nil
}
