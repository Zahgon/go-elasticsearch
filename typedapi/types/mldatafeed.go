package types

type MLDatafeed struct {
	Aggregations map[string]Aggregations `json:"aggregations,omitempty"`

	Authorization          *DatafeedAuthorization `json:"authorization,omitempty"`
	ChunkingConfig         *ChunkingConfig        `json:"chunking_config,omitempty"`
	DatafeedId             string                 `json:"datafeed_id"`
	DelayedDataCheckConfig DelayedDataCheckConfig `json:"delayed_data_check_config"`

	Frequency        Duration               `json:"frequency,omitempty"`
	Indexes          []string               `json:"indexes,omitempty"`
	Indices          []string               `json:"indices"`
	IndicesOptions   *IndicesOptions        `json:"indices_options,omitempty"`
	JobId            string                 `json:"job_id"`
	MaxEmptySearches *int                   `json:"max_empty_searches,omitempty"`
	Query            Query                  `json:"query"`
	QueryDelay       Duration               `json:"query_delay,omitempty"`
	RuntimeMappings  RuntimeFields          `json:"runtime_mappings,omitempty"`
	ScriptFields     map[string]ScriptField `json:"script_fields,omitempty"`
	ScrollSize       *int                   `json:"scroll_size,omitempty"`
}

func (s *MLDatafeed) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMLDatafeed() *MLDatafeed { _ = "STUB: not implemented"; return nil }
