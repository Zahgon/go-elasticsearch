package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _migrateAction struct {
	v *types.MigrateAction
}

func NewMigrateAction() *_migrateAction { _ = "STUB: not implemented"; return nil }

func (s *_migrateAction) Enabled(enabled bool) *_migrateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_migrateAction) MigrateActionCaster() *types.MigrateAction {
	_ = "STUB: not implemented"
	return nil
}
