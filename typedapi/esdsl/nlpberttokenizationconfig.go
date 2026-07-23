package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _nlpBertTokenizationConfig struct {
	v *types.NlpBertTokenizationConfig
}

func NewNlpBertTokenizationConfig() *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) DoLowerCase(dolowercase bool) *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) Span(span int) *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_nlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpBertTokenizationConfig) NlpBertTokenizationConfigCaster() *types.NlpBertTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
