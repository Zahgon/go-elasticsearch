package types

type CompletionToolType any

type CompletionToolTypeVariant interface {
	CompletionToolTypeCaster() *CompletionToolType
}
