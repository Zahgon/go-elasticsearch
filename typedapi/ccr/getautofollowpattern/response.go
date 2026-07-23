package getautofollowpattern

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Patterns []types.AutoFollowPattern `json:"patterns"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
