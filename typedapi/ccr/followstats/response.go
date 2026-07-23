package followstats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Indices []types.FollowIndexStats `json:"indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
