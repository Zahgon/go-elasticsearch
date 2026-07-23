package previewdatafeed

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	DatafeedConfig *types.DatafeedConfig `json:"datafeed_config,omitempty"`

	JobConfig *types.JobConfig `json:"job_config,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
