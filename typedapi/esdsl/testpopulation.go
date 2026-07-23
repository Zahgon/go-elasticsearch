package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _testPopulation struct {
	v *types.TestPopulation
}

func NewTestPopulation() *_testPopulation { _ = "STUB: not implemented"; return nil }

func (s *_testPopulation) Field(field string) *_testPopulation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_testPopulation) Filter(filter types.QueryVariant) *_testPopulation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_testPopulation) Script(script types.ScriptVariant) *_testPopulation {
	_ = "STUB: not implemented"
	return nil
}

func (s *_testPopulation) TestPopulationCaster() *types.TestPopulation {
	_ = "STUB: not implemented"
	return nil
}
