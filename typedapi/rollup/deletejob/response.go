package deletejob

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged bool                `json:"acknowledged"`
	TaskFailures []types.TaskFailure `json:"task_failures,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
