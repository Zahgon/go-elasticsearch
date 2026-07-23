package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/snowballlanguage"
)

type SnowballAnalyzer struct {
	Language  snowballlanguage.SnowballLanguage `json:"language"`
	Stopwords StopWords                         `json:"stopwords,omitempty"`
	Type      string                            `json:"type,omitempty"`
	Version   *string                           `json:"version,omitempty"`
}

func (s *SnowballAnalyzer) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func (s SnowballAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewSnowballAnalyzer() *SnowballAnalyzer { _ = "STUB: not implemented"; return nil }

type SnowballAnalyzerVariant interface {
	SnowballAnalyzerCaster() *SnowballAnalyzer
}

func (s *SnowballAnalyzer) SnowballAnalyzerCaster() *SnowballAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *SnowballAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
