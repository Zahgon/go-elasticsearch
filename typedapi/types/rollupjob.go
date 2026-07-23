package types

type RollupJob struct {
	Config RollupJobConfiguration `json:"config"`

	Stats RollupJobStats `json:"stats"`

	Status RollupJobStatus `json:"status"`
}

func NewRollupJob() *RollupJob { _ = "STUB: not implemented"; return nil }
