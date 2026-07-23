package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationmode"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/icunormalizationtype"
)

type _icuAnalyzer struct {
	v *types.IcuAnalyzer
}

func NewIcuAnalyzer(method icunormalizationtype.IcuNormalizationType, mode icunormalizationmode.IcuNormalizationMode) *_icuAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuAnalyzer) Method(method icunormalizationtype.IcuNormalizationType) *_icuAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuAnalyzer) Mode(mode icunormalizationmode.IcuNormalizationMode) *_icuAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_icuAnalyzer) IcuAnalyzerCaster() *types.IcuAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
