package getasyncstatus

type Response struct {
	CompletionStatus *uint `json:"completion_status,omitempty"`

	ExpirationTimeInMillis int64 `json:"expiration_time_in_millis"`

	Id string `json:"id"`

	IsPartial bool `json:"is_partial"`

	IsRunning bool `json:"is_running"`

	StartTimeInMillis int64 `json:"start_time_in_millis"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
