package follow

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type Request struct {
	DataStreamName *string `json:"data_stream_name,omitempty"`

	LeaderIndex string `json:"leader_index"`

	MaxOutstandingReadRequests *int64 `json:"max_outstanding_read_requests,omitempty"`

	MaxOutstandingWriteRequests *int `json:"max_outstanding_write_requests,omitempty"`

	MaxReadRequestOperationCount *int `json:"max_read_request_operation_count,omitempty"`

	MaxReadRequestSize types.ByteSize `json:"max_read_request_size,omitempty"`

	MaxRetryDelay types.Duration `json:"max_retry_delay,omitempty"`

	MaxWriteBufferCount *int `json:"max_write_buffer_count,omitempty"`

	MaxWriteBufferSize types.ByteSize `json:"max_write_buffer_size,omitempty"`

	MaxWriteRequestOperationCount *int `json:"max_write_request_operation_count,omitempty"`

	MaxWriteRequestSize types.ByteSize `json:"max_write_request_size,omitempty"`

	ReadPollTimeout types.Duration `json:"read_poll_timeout,omitempty"`

	RemoteCluster string `json:"remote_cluster"`

	Settings *types.IndexSettings `json:"settings,omitempty"`
}

func NewRequest() *Request { _ = "STUB: not implemented"; return nil }

func (r *Request) FromJSON(data string) (*Request, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *Request) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }
