package types

type Tokenizer any

type TokenizerVariant interface {
	TokenizerCaster() *Tokenizer
}
