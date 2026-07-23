package get

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Completed bool              `json:"completed"`
	Error     *types.ErrorCause `json:"error,omitempty"`
	Response  json.RawMessage   `json:"response,omitempty"`
	Task      types.TaskInfo    `json:"task"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
