package getbuckets

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Buckets []types.BucketSummary `json:"buckets"`
	Count   int64                 `json:"count"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
