package types

type Ids []string

type IdsVariant interface {
	IdsCaster() *Ids
}
