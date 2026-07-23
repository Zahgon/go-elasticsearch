package types

type SearchInputRequestBody struct {
	Query Query `json:"query"`
}

func NewSearchInputRequestBody() *SearchInputRequestBody { _ = "STUB: not implemented"; return nil }

type SearchInputRequestBodyVariant interface {
	SearchInputRequestBodyCaster() *SearchInputRequestBody
}

func (s *SearchInputRequestBody) SearchInputRequestBodyCaster() *SearchInputRequestBody {
	_ = "STUB: not implemented"
	return nil
}
