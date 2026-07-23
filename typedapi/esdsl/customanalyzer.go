package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _customAnalyzer struct {
	v *types.CustomAnalyzer
}

func NewCustomAnalyzer(tokenizer string) *_customAnalyzer { _ = "STUB: not implemented"; return nil }

func (s *_customAnalyzer) CharFilter(charfilters ...string) *_customAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customAnalyzer) Filter(filters ...string) *_customAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customAnalyzer) PositionIncrementGap(positionincrementgap int) *_customAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customAnalyzer) PositionOffsetGap(positionoffsetgap int) *_customAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customAnalyzer) Tokenizer(tokenizer string) *_customAnalyzer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_customAnalyzer) CustomAnalyzerCaster() *types.CustomAnalyzer {
	_ = "STUB: not implemented"
	return nil
}
