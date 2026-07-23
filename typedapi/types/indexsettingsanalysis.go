package types

type IndexSettingsAnalysis struct {
	Analyzer   map[string]Analyzer    `json:"analyzer,omitempty"`
	CharFilter map[string]CharFilter  `json:"char_filter,omitempty"`
	Filter     map[string]TokenFilter `json:"filter,omitempty"`
	Normalizer map[string]Normalizer  `json:"normalizer,omitempty"`
	Tokenizer  map[string]Tokenizer   `json:"tokenizer,omitempty"`
}

func (s *IndexSettingsAnalysis) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewIndexSettingsAnalysis() *IndexSettingsAnalysis { _ = "STUB: not implemented"; return nil }

type IndexSettingsAnalysisVariant interface {
	IndexSettingsAnalysisCaster() *IndexSettingsAnalysis
}

func (s *IndexSettingsAnalysis) IndexSettingsAnalysisCaster() *IndexSettingsAnalysis {
	_ = "STUB: not implemented"
	return nil
}
