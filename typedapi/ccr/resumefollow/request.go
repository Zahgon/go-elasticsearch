package resumefollow

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	MaxOutstandingReadRequests    *int64         `json:"max_outstanding_read_requests,omitempty"`
	MaxOutstandingWriteRequests   *int64         `json:"max_outstanding_write_requests,omitempty"`
	MaxReadRequestOperationCount  *int64         `json:"max_read_request_operation_count,omitempty"`
	MaxReadRequestSize            *string        `json:"max_read_request_size,omitempty"`
	MaxRetryDelay                 types.Duration `json:"max_retry_delay,omitempty"`
	MaxWriteBufferCount           *int64         `json:"max_write_buffer_count,omitempty"`
	MaxWriteBufferSize            *string        `json:"max_write_buffer_size,omitempty"`
	MaxWriteRequestOperationCount *int64         `json:"max_write_request_operation_count,omitempty"`
	MaxWriteRequestSize           *string        `json:"max_write_request_size,omitempty"`
	ReadPollTimeout               types.Duration `json:"read_poll_timeout,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
