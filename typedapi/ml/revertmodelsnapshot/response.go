package revertmodelsnapshot

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Model types.ModelSnapshot `json:"model"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
