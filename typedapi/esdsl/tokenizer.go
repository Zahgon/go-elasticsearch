package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _tokenizer struct {
	v types.Tokenizer
}

func NewTokenizer() *_tokenizer { _ = "STUB: not implemented"; return nil }

func (u *_tokenizer) String(string string) *_tokenizer { _ = "STUB: not implemented"; return nil }

func (u *_tokenizer) TokenizerDefinition(tokenizerdefinition types.TokenizerDefinitionVariant) *_tokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_tokenizerDefinition) TokenizerCaster() *types.Tokenizer {
	_ = "STUB: not implemented"
	return nil
}

func (u *_tokenizer) TokenizerCaster() *types.Tokenizer { _ = "STUB: not implemented"; return nil }
