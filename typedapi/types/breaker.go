package types

type Breaker struct {
	EstimatedSize *string `json:"estimated_size,omitempty"`

	EstimatedSizeInBytes *int64 `json:"estimated_size_in_bytes,omitempty"`

	LimitSize *string `json:"limit_size,omitempty"`

	LimitSizeInBytes *int64 `json:"limit_size_in_bytes,omitempty"`

	Overhead *float32 `json:"overhead,omitempty"`

	Tripped *float32 `json:"tripped,omitempty"`
}

func (s *Breaker) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewBreaker() *Breaker { _ = "STUB: not implemented"; return nil }
