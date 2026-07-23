package types

type DataframeAnalyticsStatsContainer struct {
	ClassificationStats *DataframeAnalyticsStatsHyperparameters `json:"classification_stats,omitempty"`

	OutlierDetectionStats *DataframeAnalyticsStatsOutlierDetection `json:"outlier_detection_stats,omitempty"`

	RegressionStats *DataframeAnalyticsStatsHyperparameters `json:"regression_stats,omitempty"`
}

func NewDataframeAnalyticsStatsContainer() *DataframeAnalyticsStatsContainer {
	_ = "STUB: not implemented"
	return nil
}
