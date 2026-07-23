package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _indexAndDataStreamAction struct {
	v *types.IndexAndDataStreamAction
}

func NewIndexAndDataStreamAction() *_indexAndDataStreamAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAndDataStreamAction) DataStream(datastreamname string) *_indexAndDataStreamAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAndDataStreamAction) Index(indexname string) *_indexAndDataStreamAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAndDataStreamAction) IndicesModifyActionCaster() *types.IndicesModifyAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_indexAndDataStreamAction) IndexAndDataStreamActionCaster() *types.IndexAndDataStreamAction {
	_ = "STUB: not implemented"
	return nil
}
