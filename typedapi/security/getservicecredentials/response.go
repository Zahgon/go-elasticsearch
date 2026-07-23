package getservicecredentials

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Count int `json:"count"`

	NodesCredentials types.NodesCredentials    `json:"nodes_credentials"`
	ServiceAccount   string                    `json:"service_account"`
	Tokens           map[string]types.Metadata `json:"tokens"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
