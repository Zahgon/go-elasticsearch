package globalcheckpoints

type Response struct {
	GlobalCheckpoints []int64 `json:"global_checkpoints"`
	TimedOut          bool    `json:"timed_out"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
