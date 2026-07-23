package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _nlpTokenizationUpdateOptions struct {
	v *types.NlpTokenizationUpdateOptions
}

func NewNlpTokenizationUpdateOptions() *_nlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpTokenizationUpdateOptions) Span(span int) *_nlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpTokenizationUpdateOptions) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpTokenizationUpdateOptions) NlpTokenizationUpdateOptionsCaster() *types.NlpTokenizationUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
