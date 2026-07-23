package types

type AnalyzeDetail struct {
	Analyzer       *AnalyzerDetail    `json:"analyzer,omitempty"`
	Charfilters    []CharFilterDetail `json:"charfilters,omitempty"`
	CustomAnalyzer bool               `json:"custom_analyzer"`
	Tokenfilters   []TokenDetail      `json:"tokenfilters,omitempty"`
	Tokenizer      *TokenDetail       `json:"tokenizer,omitempty"`
}

func (s *AnalyzeDetail) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewAnalyzeDetail() *AnalyzeDetail { _ = "STUB: not implemented"; return nil }
