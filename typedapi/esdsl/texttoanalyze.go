package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _textToAnalyze struct {
	v types.TextToAnalyze
}

func NewTextToAnalyze() *_textToAnalyze { _ = "STUB: not implemented"; return nil }

func (u *_textToAnalyze) Strings(strings ...string) *_textToAnalyze {
	_ = "STUB: not implemented"
	return nil
}

func (u *_textToAnalyze) TextToAnalyzeCaster() *types.TextToAnalyze {
	_ = "STUB: not implemented"
	return nil
}
