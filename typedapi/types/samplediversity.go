package types

type SampleDiversity struct {
	Field           string `json:"field"`
	MaxDocsPerValue int    `json:"max_docs_per_value"`
}

func (s *SampleDiversity) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewSampleDiversity() *SampleDiversity { _ = "STUB: not implemented"; return nil }

type SampleDiversityVariant interface {
	SampleDiversityCaster() *SampleDiversity
}

func (s *SampleDiversity) SampleDiversityCaster() *SampleDiversity {
	_ = "STUB: not implemented"
	return nil
}
