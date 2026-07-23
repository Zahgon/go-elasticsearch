package getnode

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Nodes []types.NodeShutdownStatus `json:"nodes"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
