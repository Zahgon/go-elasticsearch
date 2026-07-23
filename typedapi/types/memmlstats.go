package types

type MemMlStats struct {
	AnomalyDetectors ByteSize `json:"anomaly_detectors,omitempty"`

	AnomalyDetectorsInBytes int `json:"anomaly_detectors_in_bytes"`

	DataFrameAnalytics ByteSize `json:"data_frame_analytics,omitempty"`

	DataFrameAnalyticsInBytes int `json:"data_frame_analytics_in_bytes"`

	Max ByteSize `json:"max,omitempty"`

	MaxInBytes int `json:"max_in_bytes"`

	NativeCodeOverhead ByteSize `json:"native_code_overhead,omitempty"`

	NativeCodeOverheadInBytes int `json:"native_code_overhead_in_bytes"`

	NativeInference ByteSize `json:"native_inference,omitempty"`

	NativeInferenceInBytes int `json:"native_inference_in_bytes"`
}

func (s *MemMlStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMemMlStats() *MemMlStats { _ = "STUB: not implemented"; return nil }
