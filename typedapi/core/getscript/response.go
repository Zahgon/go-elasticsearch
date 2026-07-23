package getscript

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Found  bool                `json:"found"`
	Id_    string              `json:"_id"`
	Script *types.StoredScript `json:"script,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
