package types

type TokenFilter any

type TokenFilterVariant interface {
	TokenFilterCaster() *TokenFilter
}
