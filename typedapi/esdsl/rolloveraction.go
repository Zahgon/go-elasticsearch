package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _rolloverAction struct {
	v *types.RolloverAction
}

func NewRolloverAction() *_rolloverAction { _ = "STUB: not implemented"; return nil }

func (s *_rolloverAction) MaxAge(duration types.DurationVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MaxDocs(maxdocs int64) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MaxPrimaryShardDocs(maxprimarysharddocs int64) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MaxPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MaxSize(bytesize types.ByteSizeVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MinAge(duration types.DurationVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MinDocs(mindocs int64) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MinPrimaryShardDocs(minprimarysharddocs int64) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MinPrimaryShardSize(bytesize types.ByteSizeVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) MinSize(bytesize types.ByteSizeVariant) *_rolloverAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_rolloverAction) RolloverActionCaster() *types.RolloverAction {
	_ = "STUB: not implemented"
	return nil
}
