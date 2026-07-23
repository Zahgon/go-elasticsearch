package types

type IndexTemplateSummaryWithRollover struct {
	Aliases           map[string]Alias                 `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptions               `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`

	Mappings *TypeMapping `json:"mappings,omitempty"`

	Settings *IndexSettings `json:"settings,omitempty"`
}

func NewIndexTemplateSummaryWithRollover() *IndexTemplateSummaryWithRollover {
	_ = "STUB: not implemented"
	return nil
}
