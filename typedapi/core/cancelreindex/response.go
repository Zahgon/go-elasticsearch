package cancelreindex

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	Acknowledged *bool `json:"acknowledged,omitempty"`

	Cancelled *bool `json:"cancelled,omitempty"`

	Completed *bool `json:"completed,omitempty"`

	Description *string `json:"description,omitempty"`

	Error *types.ErrorCause `json:"error,omitempty"`

	Id *string `json:"id,omitempty"`

	Response *types.ReindexTaskResult `json:"response,omitempty"`

	RunningTime types.Duration `json:"running_time,omitempty"`

	RunningTimeInNanos *int64 `json:"running_time_in_nanos,omitempty"`

	StartTime *string `json:"start_time,omitempty"`

	StartTimeInMillis *int64 `json:"start_time_in_millis,omitempty"`

	Status *types.ReindexStatus `json:"status,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
