package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/noridecompoundmode"
)

type NoriAnalyzer struct {
	DecompoundMode *noridecompoundmode.NoriDecompoundMode `json:"decompound_mode,omitempty"`
	Stoptags       []string                               `json:"stoptags,omitempty"`
	Type           string                                 `json:"type,omitempty"`
	UserDictionary *string                                `json:"user_dictionary,omitempty"`
	Version        *string                                `json:"version,omitempty"`
}

func (s *NoriAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s NoriAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewNoriAnalyzer() *NoriAnalyzer { _ = "STUB: not implemented"; return nil }

type NoriAnalyzerVariant interface {
	NoriAnalyzerCaster() *NoriAnalyzer
}

func (s *NoriAnalyzer) NoriAnalyzerCaster() *NoriAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *NoriAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
