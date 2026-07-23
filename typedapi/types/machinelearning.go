package types

type MachineLearning struct {
	Available              bool                     `json:"available"`
	DataFrameAnalyticsJobs MlDataFrameAnalyticsJobs `json:"data_frame_analytics_jobs"`
	Datafeeds              map[string]XpackDatafeed `json:"datafeeds"`
	Enabled                bool                     `json:"enabled"`
	Inference              MlInference              `json:"inference"`

	Jobs      map[string]JobUsage `json:"jobs"`
	NodeCount int                 `json:"node_count"`
}

func (s *MachineLearning) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMachineLearning() *MachineLearning { _ = "STUB: not implemented"; return nil }
