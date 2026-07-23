package types

type SnapshotNodeInfo struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

func (s *SnapshotNodeInfo) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSnapshotNodeInfo() *SnapshotNodeInfo { _ = "STUB: not implemented"; return nil }
