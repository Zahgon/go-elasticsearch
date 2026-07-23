package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _xlmRobertaTokenizationConfig struct {
	v *types.XlmRobertaTokenizationConfig
}

func NewXlmRobertaTokenizationConfig() *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) DoLowerCase(dolowercase bool) *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) Span(span int) *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_xlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_xlmRobertaTokenizationConfig) XlmRobertaTokenizationConfigCaster() *types.XlmRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
