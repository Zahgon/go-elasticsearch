package activatewatch

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Status types.ActivationStatus `json:"status"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
