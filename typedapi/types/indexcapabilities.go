package types

type IndexCapabilities struct {
	RollupJobs []RollupJobSummary `json:"rollup_jobs"`
}

func NewIndexCapabilities() *IndexCapabilities { _ = "STUB: not implemented"; return nil }
