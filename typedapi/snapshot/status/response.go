package status

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Snapshots []types.Status `json:"snapshots"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
