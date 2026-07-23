package types

type TokenFilterDefinition any

type TokenFilterDefinitionVariant interface {
	TokenFilterDefinitionCaster() *TokenFilterDefinition
}
