package types

type Invocation struct {
	SnapshotName string   `json:"snapshot_name"`
	Time         DateTime `json:"time"`
}

func (s *Invocation) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewInvocation() *Invocation { _ = "STUB: not implemented"; return nil }
