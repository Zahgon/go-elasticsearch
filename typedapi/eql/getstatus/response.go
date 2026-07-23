package getstatus

type Response struct {
	CompletionStatus *int `json:"completion_status,omitempty"`

	ExpirationTimeInMillis *int64 `json:"expiration_time_in_millis,omitempty"`

	Id string `json:"id"`

	IsPartial bool `json:"is_partial"`

	IsRunning bool `json:"is_running"`

	StartTimeInMillis *int64 `json:"start_time_in_millis,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
