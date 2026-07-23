package types

type RollupJobSummary struct {
	Fields       map[string][]RollupJobSummaryField `json:"fields"`
	IndexPattern string                             `json:"index_pattern"`
	JobId        string                             `json:"job_id"`
	RollupIndex  string                             `json:"rollup_index"`
}

func (s *RollupJobSummary) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewRollupJobSummary() *RollupJobSummary { _ = "STUB: not implemented"; return nil }
