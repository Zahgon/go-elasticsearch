package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _intervalsRange struct {
	v *types.IntervalsRange
}

func NewIntervalsRange() *_intervalsRange { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) Analyzer(analyzer string) *_intervalsRange {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRange) Gt(gt string) *_intervalsRange { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) Gte(gte string) *_intervalsRange { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) Lt(lt string) *_intervalsRange { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) Lte(lte string) *_intervalsRange { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) UseField(field string) *_intervalsRange {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRange) IntervalsCaster() *types.Intervals { _ = "STUB: not implemented"; return nil }

func (s *_intervalsRange) IntervalsQueryCaster() *types.IntervalsQuery {
	_ = "STUB: not implemented"
	return nil
}

func (s *_intervalsRange) IntervalsRangeCaster() *types.IntervalsRange {
	_ = "STUB: not implemented"
	return nil
}
