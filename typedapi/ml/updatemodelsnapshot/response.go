package updatemodelsnapshot

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged bool                `json:"acknowledged"`
	Model        types.ModelSnapshot `json:"model"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
