package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _downsamplingRound struct {
	v *types.DownsamplingRound
}

func NewDownsamplingRound() *_downsamplingRound { _ = "STUB: not implemented"; return nil }

func (s *_downsamplingRound) After(duration types.DurationVariant) *_downsamplingRound {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsamplingRound) FixedInterval(durationlarge string) *_downsamplingRound {
	_ = "STUB: not implemented"
	return nil
}

func (s *_downsamplingRound) DownsamplingRoundCaster() *types.DownsamplingRound {
	_ = "STUB: not implemented"
	return nil
}
