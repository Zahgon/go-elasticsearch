package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _downsampleAction struct {
	v *types.DownsampleAction
}

func NewDownsampleAction() *_downsampleAction { _ = "STUB: not implemented"; return nil }

func (s *_downsampleAction) FixedInterval(durationlarge string) *_downsampleAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsampleAction) WaitTimeout(duration types.DurationVariant) *_downsampleAction {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsampleAction) DownsampleActionCaster() *types.DownsampleAction {
	_ = "STUB: not implemented"
	return nil
}
