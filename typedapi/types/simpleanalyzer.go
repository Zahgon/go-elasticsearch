package types

type SimpleAnalyzer struct {
	Type    string  `json:"type,omitempty"`
	Version *string `json:"version,omitempty"`
}

func (s *SimpleAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SimpleAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSimpleAnalyzer() *SimpleAnalyzer { _ = "STUB: not implemented"; return nil }

type SimpleAnalyzerVariant interface {
	SimpleAnalyzerCaster() *SimpleAnalyzer
}

func (s *SimpleAnalyzer) SimpleAnalyzerCaster() *SimpleAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SimpleAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
