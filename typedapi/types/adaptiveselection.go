package types

type AdaptiveSelection struct {
	AvgQueueSize *int64 `json:"avg_queue_size,omitempty"`

	AvgResponseTime Duration `json:"avg_response_time,omitempty"`

	AvgResponseTimeNs *int64 `json:"avg_response_time_ns,omitempty"`

	AvgServiceTime Duration `json:"avg_service_time,omitempty"`

	AvgServiceTimeNs *int64 `json:"avg_service_time_ns,omitempty"`

	OutgoingSearches *int64 `json:"outgoing_searches,omitempty"`

	Rank *string `json:"rank,omitempty"`
}

func (s *AdaptiveSelection) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAdaptiveSelection() *AdaptiveSelection { _ = "STUB: not implemented"; return nil }
