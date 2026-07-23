package types

type JvmStats struct {
	HeapMax ByteSize `json:"heap_max,omitempty"`

	HeapMaxInBytes int `json:"heap_max_in_bytes"`

	JavaInference ByteSize `json:"java_inference,omitempty"`

	JavaInferenceInBytes int `json:"java_inference_in_bytes"`

	JavaInferenceMax ByteSize `json:"java_inference_max,omitempty"`

	JavaInferenceMaxInBytes int `json:"java_inference_max_in_bytes"`
}

func (s *JvmStats) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewJvmStats() *JvmStats { _ = "STUB: not implemented"; return nil }
