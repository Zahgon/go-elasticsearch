package types

type MinimumShouldMatch any

type MinimumShouldMatchVariant interface {
	MinimumShouldMatchCaster() *MinimumShouldMatch
}
