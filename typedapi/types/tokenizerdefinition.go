package types

type TokenizerDefinition any

type TokenizerDefinitionVariant interface {
	TokenizerDefinitionCaster() *TokenizerDefinition
}
