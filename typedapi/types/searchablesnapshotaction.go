package types

type SearchableSnapshotAction struct {
	ForceMergeIndex    *bool  `json:"force_merge_index,omitempty"`
	SnapshotRepository string `json:"snapshot_repository"`
}

func (s *SearchableSnapshotAction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewSearchableSnapshotAction() *SearchableSnapshotAction { _ = "STUB: not implemented"; return nil }

type SearchableSnapshotActionVariant interface {
	SearchableSnapshotActionCaster() *SearchableSnapshotAction
}

func (s *SearchableSnapshotAction) SearchableSnapshotActionCaster() *SearchableSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}
