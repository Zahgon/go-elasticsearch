package getoverallbuckets

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int64 `json:"count"`

	OverallBuckets []types.OverallBucket `json:"overall_buckets"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
