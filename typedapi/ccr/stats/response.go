package stats

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	AutoFollowStats types.AutoFollowStats `json:"auto_follow_stats"`

	FollowStats types.FollowStats `json:"follow_stats"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
