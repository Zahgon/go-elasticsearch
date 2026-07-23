package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _deleteAction struct {
	v *types.DeleteAction
}

func NewDeleteAction() *_deleteAction { _ = "STUB: not implemented"; return nil }

func (s *_deleteAction) DeleteSearchableSnapshot(deletesearchablesnapshot bool) *_deleteAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_deleteAction) DeleteActionCaster() *types.DeleteAction {
	_ = "STUB: not implemented"
	return nil
}
