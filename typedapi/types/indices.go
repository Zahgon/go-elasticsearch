package types

type Indices []string

type IndicesVariant interface {
	IndicesCaster() *Indices
}
