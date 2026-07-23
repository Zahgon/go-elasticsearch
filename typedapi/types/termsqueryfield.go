package types

type TermsQueryField any

type TermsQueryFieldVariant interface {
	TermsQueryFieldCaster() *TermsQueryField
}
