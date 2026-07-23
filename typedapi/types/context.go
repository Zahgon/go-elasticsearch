package types

type Context any

type ContextVariant interface {
	ContextCaster() *Context
}
