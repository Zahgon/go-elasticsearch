package pendingtasks

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Tasks []types.PendingTask `json:"tasks"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
