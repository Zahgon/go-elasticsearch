package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _shrinkAction struct {
	v *types.ShrinkAction
}

func NewShrinkAction() *_shrinkAction { _ = "STUB: not implemented"; return nil }

func (s *_shrinkAction) AllowWriteAfterShrink(allowwriteaftershrink bool) *_shrinkAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shrinkAction) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_shrinkAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shrinkAction) NumberOfShards(numberofshards int) *_shrinkAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_shrinkAction) ShrinkActionCaster() *types.ShrinkAction {
	_ = "STUB: not implemented"
	return nil
}
