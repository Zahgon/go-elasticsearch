package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/optype"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/refresh"
)

type _indexAction struct {
	v *types.IndexAction
}

func NewIndexAction() *_indexAction { _ = "STUB: not implemented"; return nil }

func (s *_indexAction) DocId(id string) *_indexAction { _ = "STUB: not implemented"; return nil }

func (s *_indexAction) ExecutionTimeField(field string) *_indexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAction) Index(indexname string) *_indexAction { _ = "STUB: not implemented"; return nil }

func (s *_indexAction) OpType(optype optype.OpType) *_indexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAction) Refresh(refresh refresh.Refresh) *_indexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAction) Timeout(duration types.DurationVariant) *_indexAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAction) IndexActionCaster() *types.IndexAction {
	_ = "STUB: not implemented"
	return nil
}
