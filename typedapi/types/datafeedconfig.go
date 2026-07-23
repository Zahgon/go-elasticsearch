package types

type DatafeedConfig struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`

	ChunkingConfig *ChunkingConfig `json:"chunking_config,omitempty"`

	DatafeedId *string `json:"datafeed_id,omitempty"`

	DelayedDataCheckConfig *DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`

	Frequency Duration `json:"frequency,omitempty"`

	Indices []string `json:"indices,omitempty"`

	IndicesOptions *IndicesOptions `json:"indices_options,omitempty"`
	JobId          *string         `json:"job_id,omitempty"`

	MaxEmptySearches *int `json:"max_empty_searches,omitempty"`

	Query *Query `json:"query,omitempty"`

	QueryDelay Duration `json:"query_delay,omitempty"`

	RuntimeMappings RuntimeFields `json:"runtime_mappings,omitempty"`

	ScriptFields map[string]ScriptField `json:"script_fields,omitempty"`

	ScrollSize *int `json:"scroll_size,omitempty"`
}

func (s *DatafeedConfig) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewDatafeedConfig() *DatafeedConfig { _ = "STUB: not implemented"; return nil }

type DatafeedConfigVariant interface {
	DatafeedConfigCaster() *DatafeedConfig
}

func (s *DatafeedConfig) DatafeedConfigCaster() *DatafeedConfig {
	_ = "STUB: not implemented"
	return nil
}
