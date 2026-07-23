package types

type RollupCapabilities struct {
	RollupJobs []RollupCapabilitySummary `json:"rollup_jobs"`
}

func NewRollupCapabilities() *RollupCapabilities { _ = "STUB: not implemented"; return nil }
