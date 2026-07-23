package types

type TrainedModelDeploymentNodesStats struct {
	AverageInferenceProcessMemoryRssBytes ByteSize `json:"average_inference_process_memory_rss_bytes,omitempty"`

	AverageInferenceTimeMs Float64 `json:"average_inference_time_ms,omitempty"`

	AverageInferenceTimeMsExcludingCacheHits Float64 `json:"average_inference_time_ms_excluding_cache_hits,omitempty"`
	AverageInferenceTimeMsLastMinute         Float64 `json:"average_inference_time_ms_last_minute,omitempty"`

	ErrorCount                       *int   `json:"error_count,omitempty"`
	InferenceCacheHitCount           *int64 `json:"inference_cache_hit_count,omitempty"`
	InferenceCacheHitCountLastMinute *int64 `json:"inference_cache_hit_count_last_minute,omitempty"`

	InferenceCount *int64 `json:"inference_count,omitempty"`

	LastAccess *int64 `json:"last_access,omitempty"`

	Node DiscoveryNode `json:"node,omitempty"`

	NumberOfAllocations *int `json:"number_of_allocations,omitempty"`

	NumberOfPendingRequests *int  `json:"number_of_pending_requests,omitempty"`
	PeakThroughputPerMinute int64 `json:"peak_throughput_per_minute"`

	RejectedExecutionCount *int `json:"rejected_execution_count,omitempty"`

	RoutingState TrainedModelAssignmentRoutingStateAndReason `json:"routing_state"`

	StartTime *int64 `json:"start_time,omitempty"`

	ThreadsPerAllocation *int `json:"threads_per_allocation,omitempty"`
	ThroughputLastMinute int  `json:"throughput_last_minute"`

	TimeoutCount *int `json:"timeout_count,omitempty"`
}

func (s *TrainedModelDeploymentNodesStats) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTrainedModelDeploymentNodesStats() *TrainedModelDeploymentNodesStats {
	_ = "STUB: not implemented"
	return nil
}
