package esdsl

import (
	"encoding/json"

	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
)

type _updateAction struct {
	v *types.UpdateAction
}

func NewUpdateAction() *_updateAction { _ = "STUB: not implemented"; return nil }

func (s *_updateAction) DetectNoop(detectnoop bool) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) Doc(doc json.RawMessage) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) DocAsUpsert(docasupsert bool) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) Script(script types.ScriptVariant) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) ScriptedUpsert(scriptedupsert bool) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) Source_(sourceconfig types.SourceConfigVariant) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) Upsert(upsert json.RawMessage) *_updateAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_updateAction) UpdateActionCaster() *types.UpdateAction {
	_ = "STUB: not implemented"
	return nil
}
