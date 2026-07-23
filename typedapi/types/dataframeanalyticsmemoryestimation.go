package types

type DataframeAnalyticsMemoryEstimation struct {
	ExpectedMemoryWithDisk string `json:"expected_memory_with_disk"`

	ExpectedMemoryWithoutDisk string `json:"expected_memory_without_disk"`
}

func (s *DataframeAnalyticsMemoryEstimation) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsMemoryEstimation() *DataframeAnalyticsMemoryEstimation {
	_ = "STUB: not implemented"
	return nil
}
