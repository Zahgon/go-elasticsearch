package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _forceMergeAction struct {
	v *types.ForceMergeAction
}

func NewForceMergeAction(maxnumsegments int) *_forceMergeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_forceMergeAction) IndexCodec(indexcodec string) *_forceMergeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_forceMergeAction) MaxNumSegments(maxnumsegments int) *_forceMergeAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_forceMergeAction) ForceMergeActionCaster() *types.ForceMergeAction {
	_ = "STUB: not implemented"
	return nil
}
