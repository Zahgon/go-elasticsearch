package submit

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Response struct {
	CompletionTime         types.DateTime    `json:"completion_time,omitempty"`
	CompletionTimeInMillis *int64            `json:"completion_time_in_millis,omitempty"`
	Error                  *types.ErrorCause `json:"error,omitempty"`

	ExpirationTime         types.DateTime `json:"expiration_time,omitempty"`
	ExpirationTimeInMillis int64          `json:"expiration_time_in_millis"`
	Id                     *string        `json:"id,omitempty"`

	IsPartial bool `json:"is_partial"`

	IsRunning         bool              `json:"is_running"`
	Response          types.AsyncSearch `json:"response"`
	StartTime         types.DateTime    `json:"start_time,omitempty"`
	StartTimeInMillis int64             `json:"start_time_in_millis"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }

func (s *Response) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
