package types

type Merge struct {
	Scheduler *MergeScheduler `json:"scheduler,omitempty"`
}

func NewMerge() *Merge { _ = "STUB: not implemented"; return nil }

type MergeVariant interface {
	MergeCaster() *Merge
}

func (s *Merge) MergeCaster() *Merge { _ = "STUB: not implemented"; return nil }
