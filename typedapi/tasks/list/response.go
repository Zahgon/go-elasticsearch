package list

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	NodeFailures []types.ErrorCause `json:"node_failures,omitempty"`

	Nodes        map[string]types.NodeTasks `json:"nodes,omitempty"`
	TaskFailures []types.TaskFailure        `json:"task_failures,omitempty"`

	Tasks types.TaskInfos `json:"tasks,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
