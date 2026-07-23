package types

type RuntimeFieldFetchFields struct {
	Field  string  `json:"field"`
	Format *string `json:"format,omitempty"`
}

func (s *RuntimeFieldFetchFields) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewRuntimeFieldFetchFields() *RuntimeFieldFetchFields { _ = "STUB: not implemented"; return nil }

type RuntimeFieldFetchFieldsVariant interface {
	RuntimeFieldFetchFieldsCaster() *RuntimeFieldFetchFields
}

func (s *RuntimeFieldFetchFields) RuntimeFieldFetchFieldsCaster() *RuntimeFieldFetchFields {
	_ = "STUB: not implemented"
	return nil
}
