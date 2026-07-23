package getautoscalingcapacity

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Policies map[string]types.AutoscalingDeciders `json:"policies"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
