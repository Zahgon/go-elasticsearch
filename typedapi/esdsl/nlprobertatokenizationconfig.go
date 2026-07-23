package esdsl

import (
	"github.com/elastic/go-elasticsearch/v9/typedapi/types"
	"github.com/elastic/go-elasticsearch/v9/typedapi/types/enums/tokenizationtruncate"
)

type _nlpRobertaTokenizationConfig struct {
	v *types.NlpRobertaTokenizationConfig
}

func NewNlpRobertaTokenizationConfig() *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) AddPrefixSpace(addprefixspace bool) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) DoLowerCase(dolowercase bool) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) MaxSequenceLength(maxsequencelength int) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) Span(span int) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) Truncate(truncate tokenizationtruncate.TokenizationTruncate) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) WithSpecialTokens(withspecialtokens bool) *_nlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_nlpRobertaTokenizationConfig) NlpRobertaTokenizationConfigCaster() *types.NlpRobertaTokenizationConfig {
	_ = "STUB: not implemented"
	return nil
}
