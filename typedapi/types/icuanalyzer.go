package types

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type IcuAnalyzer struct {
	Method icunormalizationtype.IcuNormalizationType `json:"method"`
	Mode   icunormalizationmode.IcuNormalizationMode `json:"mode"`
	Type   string                                    `json:"type,omitempty"`
}

func (s IcuAnalyzer) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func NewIcuAnalyzer() *IcuAnalyzer { _ = "STUB: not implemented"; return nil }

type IcuAnalyzerVariant interface {
	IcuAnalyzerCaster() *IcuAnalyzer
}

func (s *IcuAnalyzer) IcuAnalyzerCaster() *IcuAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *IcuAnalyzer) AnalyzerCaster() *Analyzer { _ = "STUB: not implemented"; return nil }
