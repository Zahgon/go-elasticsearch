package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsFilter struct {
	v *types.IntervalsFilter
}

func NewIntervalsFilter() *_intervalsFilter { _ = "STUB: not implemented"; return nil }

func (s *_intervalsFilter) After(after types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) Before(before types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) ContainedBy(containedby types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) Containing(containing types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) NotContainedBy(notcontainedby types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) NotContaining(notcontaining types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) NotOverlapping(notoverlapping types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) Overlapping(overlapping types.IntervalsVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) Script(script types.ScriptVariant) *_intervalsFilter {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsFilter) IntervalsFilterCaster() *types.IntervalsFilter {
	_ = "STUB: not implemented"
	return nil
}
