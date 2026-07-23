package types

type FollowerIndexParameters struct {
	MaxOutstandingReadRequests *int64 `json:"max_outstanding_read_requests,omitempty"`

	MaxOutstandingWriteRequests *int `json:"max_outstanding_write_requests,omitempty"`

	MaxReadRequestOperationCount *int `json:"max_read_request_operation_count,omitempty"`

	MaxReadRequestSize ByteSize `json:"max_read_request_size,omitempty"`

	MaxRetryDelay Duration `json:"max_retry_delay,omitempty"`

	MaxWriteBufferCount *int `json:"max_write_buffer_count,omitempty"`

	MaxWriteBufferSize ByteSize `json:"max_write_buffer_size,omitempty"`

	MaxWriteRequestOperationCount *int `json:"max_write_request_operation_count,omitempty"`

	MaxWriteRequestSize ByteSize `json:"max_write_request_size,omitempty"`

	ReadPollTimeout Duration `json:"read_poll_timeout,omitempty"`
}

func (s *FollowerIndexParameters) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewFollowerIndexParameters() *FollowerIndexParameters { _ = "STUB: not implemented"; return nil }
