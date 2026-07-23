package putjob

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	Cron string `json:"cron"`

	Groups  types.Groupings   `json:"groups"`
	Headers types.HttpHeaders `json:"headers,omitempty"`

	IndexPattern string `json:"index_pattern"`

	Metrics []types.FieldMetric `json:"metrics,omitempty"`

	PageSize int `json:"page_size"`

	RollupIndex string `json:"rollup_index"`

	Timeout types.Duration `json:"timeout,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
