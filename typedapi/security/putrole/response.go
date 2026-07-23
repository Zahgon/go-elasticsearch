package putrole

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Role types.CreatedStatus `json:"role"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
