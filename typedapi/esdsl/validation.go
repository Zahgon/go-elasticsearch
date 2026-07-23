package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _validation struct {
	v types.Validation
}

func NewValidation() *_validation { _ = "STUB: not implemented"; return nil }

func (u *_validation) LessThanValidation(lessthanvalidation types.LessThanValidationVariant) *_validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_lessThanValidation) ValidationCaster() *types.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_validation) GreaterThanValidation(greaterthanvalidation types.GreaterThanValidationVariant) *_validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_greaterThanValidation) ValidationCaster() *types.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_validation) ListTypeValidation(listtypevalidation types.ListTypeValidationVariant) *_validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_listTypeValidation) ValidationCaster() *types.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_validation) IncludedInValidation(includedinvalidation types.IncludedInValidationVariant) *_validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_includedInValidation) ValidationCaster() *types.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_validation) RegexValidation(regexvalidation types.RegexValidationVariant) *_validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_regexValidation) ValidationCaster() *types.Validation {
	_ = "STUB: not implemented"
	return nil
}

func (u *_validation) ValidationCaster() *types.Validation { _ = "STUB: not implemented"; return nil }
