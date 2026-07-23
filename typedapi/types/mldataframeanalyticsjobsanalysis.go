package types

type MlDataFrameAnalyticsJobsAnalysis struct {
	Classification   *int `json:"classification,omitempty"`
	OutlierDetection *int `json:"outlier_detection,omitempty"`
	Regression       *int `json:"regression,omitempty"`
}

func (s *MlDataFrameAnalyticsJobsAnalysis) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewMlDataFrameAnalyticsJobsAnalysis() *MlDataFrameAnalyticsJobsAnalysis {
	_ = "STUB: not implemented"
	return nil
}
