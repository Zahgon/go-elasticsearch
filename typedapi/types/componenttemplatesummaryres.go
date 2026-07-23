package types

type ComponentTemplateSummaryRes struct {
	Aliases           map[string]AliasDefinition       `json:"aliases,omitempty"`
	DataStreamOptions *DataStreamOptions               `json:"data_stream_options,omitempty"`
	Lifecycle         *DataStreamLifecycleWithRollover `json:"lifecycle,omitempty"`
	Mappings          *TypeMapping                     `json:"mappings,omitempty"`
	Meta_             Metadata                         `json:"_meta,omitempty"`
	Settings          map[string]IndexSettings         `json:"settings,omitempty"`
	Version           *int64                           `json:"version,omitempty"`
}

func (s *ComponentTemplateSummaryRes) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewComponentTemplateSummaryRes() *ComponentTemplateSummaryRes {
	_ = "STUB: not implemented"
	return nil
}
