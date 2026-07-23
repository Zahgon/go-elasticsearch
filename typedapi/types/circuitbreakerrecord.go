package types

type CircuitBreakerRecord struct {
	Breaker *string `json:"breaker,omitempty"`

	Estimated *string `json:"estimated,omitempty"`

	EstimatedBytes ByteSize `json:"estimated_bytes,omitempty"`

	Limit *string `json:"limit,omitempty"`

	LimitBytes ByteSize `json:"limit_bytes,omitempty"`

	NodeId *string `json:"node_id,omitempty"`

	NodeName *string `json:"node_name,omitempty"`

	Overhead *string `json:"overhead,omitempty"`

	Tripped *string `json:"tripped,omitempty"`
}

func (s *CircuitBreakerRecord) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCircuitBreakerRecord() *CircuitBreakerRecord { _ = "STUB: not implemented"; return nil }
