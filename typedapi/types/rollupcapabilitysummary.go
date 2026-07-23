package types

type RollupCapabilitySummary struct {
	Fields       map[string][]RollupFieldSummary `json:"fields"`
	IndexPattern string                          `json:"index_pattern"`
	JobId        string                          `json:"job_id"`
	RollupIndex  string                          `json:"rollup_index"`
}

func (s *RollupCapabilitySummary) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRollupCapabilitySummary() *RollupCapabilitySummary { _ = "STUB: not implemented"; return nil }
