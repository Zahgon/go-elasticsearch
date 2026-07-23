package getjobs

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int64       `json:"count"`
	Jobs  []types.Job `json:"jobs"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
