package flushjob

type Response struct {
	Flushed bool `json:"flushed"`

	LastFinalizedBucketEnd *int `json:"last_finalized_bucket_end,omitempty"`
}

func NewResponse() *Response { _ = "STUB: not implemented"; return nil }
