package types

type TermsExclude []string

type TermsExcludeVariant interface {
	TermsExcludeCaster() *TermsExclude
}
