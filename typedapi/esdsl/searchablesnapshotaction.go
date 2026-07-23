package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _searchableSnapshotAction struct {
	v *types.SearchableSnapshotAction
}

func NewSearchableSnapshotAction(snapshotrepository string) *_searchableSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchableSnapshotAction) ForceMergeIndex(forcemergeindex bool) *_searchableSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchableSnapshotAction) SnapshotRepository(snapshotrepository string) *_searchableSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_searchableSnapshotAction) SearchableSnapshotActionCaster() *types.SearchableSnapshotAction {
	_ = "STUB: not implemented"
	return nil
}
