package types

type MlDataFrameAnalyticsJobs struct {
	All_           MlDataFrameAnalyticsJobsCount     `json:"_all"`
	AnalysisCounts *MlDataFrameAnalyticsJobsAnalysis `json:"analysis_counts,omitempty"`
	MemoryUsage    *MlDataFrameAnalyticsJobsMemory   `json:"memory_usage,omitempty"`
	Stopped        *MlDataFrameAnalyticsJobsCount    `json:"stopped,omitempty"`
}

func NewMlDataFrameAnalyticsJobs() *MlDataFrameAnalyticsJobs { _ = "STUB: not implemented"; return nil }
