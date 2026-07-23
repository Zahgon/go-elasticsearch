package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _tokenizationConfigContainer struct {
	v *types.TokenizationConfigContainer
}

func NewTokenizationConfigContainer() *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) Bert(bert types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) BertJa(bertja types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) Mpnet(mpnet types.NlpBertTokenizationConfigVariant) *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) Roberta(roberta types.NlpRobertaTokenizationConfigVariant) *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) XlmRoberta(xlmroberta types.XlmRobertaTokenizationConfigVariant) *_tokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_tokenizationConfigContainer) TokenizationConfigContainerCaster() *types.TokenizationConfigContainer {
	_ = "STUB: not implemented"
	return nil
}
