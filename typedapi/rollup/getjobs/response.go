package getjobs

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Jobs []types.RollupJob `json:"jobs"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
