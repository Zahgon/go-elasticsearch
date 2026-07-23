package types

type Stringifiedinteger any

type StringifiedintegerVariant interface {
	StringifiedintegerCaster() *Stringifiedinteger
}
