package types

type CharFilterDefinition any

type CharFilterDefinitionVariant interface {
	CharFilterDefinitionCaster() *CharFilterDefinition
}
