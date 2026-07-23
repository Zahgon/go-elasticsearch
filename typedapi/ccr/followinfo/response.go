package followinfo

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	FollowerIndices []types.FollowerIndex `json:"follower_indices"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
