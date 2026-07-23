package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _sampleDiversity struct {
	v *types.SampleDiversity
}

func NewSampleDiversity(maxdocspervalue int) *_sampleDiversity {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sampleDiversity) Field(field string) *_sampleDiversity {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sampleDiversity) MaxDocsPerValue(maxdocspervalue int) *_sampleDiversity {
	_ = "STUB: not implemented"
	return nil
}

func (s *_sampleDiversity) SampleDiversityCaster() *types.SampleDiversity {
	_ = "STUB: not implemented"
	return nil
}
