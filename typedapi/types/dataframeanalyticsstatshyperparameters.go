package types

type DataframeAnalyticsStatsHyperparameters struct {
	Hyperparameters Hyperparameters `json:"hyperparameters"`

	Iteration int `json:"iteration"`

	Timestamp int64 `json:"timestamp"`

	TimingStats TimingStats `json:"timing_stats"`

	ValidationLoss ValidationLoss `json:"validation_loss"`
}

func (s *DataframeAnalyticsStatsHyperparameters) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewDataframeAnalyticsStatsHyperparameters() *DataframeAnalyticsStatsHyperparameters {
	_ = "STUB: not implemented"
	return nil
}
