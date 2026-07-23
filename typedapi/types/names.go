package types

type Names []string

type NamesVariant interface {
	NamesCaster() *Names
}
