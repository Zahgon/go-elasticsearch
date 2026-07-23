package types

type Client struct {
	Agent *string `json:"agent,omitempty"`

	ClosedTimeMillis *int64 `json:"closed_time_millis,omitempty"`

	Id *int64 `json:"id,omitempty"`

	LastRequestTimeMillis *int64 `json:"last_request_time_millis,omitempty"`

	LastUri *string `json:"last_uri,omitempty"`

	LocalAddress *string `json:"local_address,omitempty"`

	OpenedTimeMillis *int64 `json:"opened_time_millis,omitempty"`

	RemoteAddress *string `json:"remote_address,omitempty"`

	RequestCount *int64 `json:"request_count,omitempty"`

	RequestSizeBytes *int64 `json:"request_size_bytes,omitempty"`

	XOpaqueId *string `json:"x_opaque_id,omitempty"`
}

func (s *Client) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewClient() *Client { _ = "STUB: not implemented"; return nil }
