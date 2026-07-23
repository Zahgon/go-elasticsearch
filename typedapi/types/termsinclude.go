package types

type TermsInclude any

type TermsIncludeVariant interface {
	TermsIncludeCaster() *TermsInclude
}
