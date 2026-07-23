package types

type ShardStatistics struct {
	Failed   uint           `json:"failed"`
	Failures []ShardFailure `json:"failures,omitempty"`
	Skipped  *uint          `json:"skipped,omitempty"`

	Successful uint `json:"successful"`

	Total uint `json:"total"`
}

func NewShardStatistics() *ShardStatistics { _ = "STUB: not implemented"; return nil }
