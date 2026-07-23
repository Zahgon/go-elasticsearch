package getoverallbuckets

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	AllowNoMatch *bool `json:"allow_no_match,omitempty"`

	BucketSpan types.Duration `json:"bucket_span,omitempty"`

	End types.DateTime `json:"end,omitempty"`

	ExcludeInterim *bool `json:"exclude_interim,omitempty"`

	OverallScore *types.Float64 `json:"overall_score,omitempty"`

	Start types.DateTime `json:"start,omitempty"`

	TopN *int `json:"top_n,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
