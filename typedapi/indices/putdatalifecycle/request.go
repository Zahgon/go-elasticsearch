package putdatalifecycle

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/samplingmethod"
)

type Request struct {
	DataRetention types.Duration `json:"data_retention,omitempty"`

	Downsampling []types.DownsamplingRound `json:"downsampling,omitempty"`

	DownsamplingMethod *samplingmethod.SamplingMethod `json:"downsampling_method,omitempty"`

	Enabled *bool `json:"enabled,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
