package types

type MergeScheduler struct {
	MaxMergeCount  Stringifiedinteger `json:"max_merge_count,omitempty"`
	MaxThreadCount Stringifiedinteger `json:"max_thread_count,omitempty"`
}

func (s *MergeScheduler) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewMergeScheduler() *MergeScheduler { _ = "STUB: not implemented"; return nil }

type MergeSchedulerVariant interface {
	MergeSchedulerCaster() *MergeScheduler
}

func (s *MergeScheduler) MergeSchedulerCaster() *MergeScheduler {
	_ = "STUB: not implemented"
	return nil
}
