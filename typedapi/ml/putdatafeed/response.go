package putdatafeed

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Aggregations           map[string]types.Aggregations `json:"aggregations,omitempty"`
	Authorization          *types.DatafeedAuthorization  `json:"authorization,omitempty"`
	ChunkingConfig         types.ChunkingConfig          `json:"chunking_config"`
	DatafeedId             string                        `json:"datafeed_id"`
	DelayedDataCheckConfig *types.DelayedDataCheckConfig `json:"delayed_data_check_config,omitempty"`
	Frequency              types.Duration                `json:"frequency,omitempty"`
	Indices                []string                      `json:"indices"`
	IndicesOptions         *types.IndicesOptions         `json:"indices_options,omitempty"`
	JobId                  string                        `json:"job_id"`
	MaxEmptySearches       *int                          `json:"max_empty_searches,omitempty"`
	Query                  types.Query                   `json:"query"`
	QueryDelay             types.Duration                `json:"query_delay"`
	RuntimeMappings        types.RuntimeFields           `json:"runtime_mappings,omitempty"`
	ScriptFields           map[string]types.ScriptField  `json:"script_fields,omitempty"`
	ScrollSize             int                           `json:"scroll_size"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
