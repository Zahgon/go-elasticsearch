package listreindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	NodeFailures []types.ErrorCause `json:"node_failures,omitempty"`

	Reindex []types.ReindexTaskInfo `json:"reindex"`

	TaskFailures []types.TaskFailure `json:"task_failures,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
