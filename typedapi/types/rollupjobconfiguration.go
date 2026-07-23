package types

type RollupJobConfiguration struct {
	Cron         string        `json:"cron"`
	Groups       Groupings     `json:"groups"`
	Id           string        `json:"id"`
	IndexPattern string        `json:"index_pattern"`
	Metrics      []FieldMetric `json:"metrics"`
	PageSize     int64         `json:"page_size"`
	RollupIndex  string        `json:"rollup_index"`
	Timeout      Duration      `json:"timeout"`
}

func (s *RollupJobConfiguration) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRollupJobConfiguration() *RollupJobConfiguration { _ = "STUB: not implemented"; return nil }
