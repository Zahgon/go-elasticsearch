package types

type SnapshotResponseItem struct {
	Error      *ErrorCause    `json:"error,omitempty"`
	Repository string         `json:"repository"`
	Snapshots  []SnapshotInfo `json:"snapshots,omitempty"`
}

func (s *SnapshotResponseItem) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSnapshotResponseItem() *SnapshotResponseItem { _ = "STUB: not implemented"; return nil }
