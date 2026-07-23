package types

type CompletionContext struct {
	Boost *Float64 `json:"boost,omitempty"`

	Context Context `json:"context"`

	Neighbours []GeoHashPrecision `json:"neighbours,omitempty"`

	Precision GeoHashPrecision `json:"precision,omitempty"`

	Prefix *bool `json:"prefix,omitempty"`
}

func (s *CompletionContext) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewCompletionContext() *CompletionContext { _ = "STUB: not implemented"; return nil }

type CompletionContextVariant interface {
	CompletionContextCaster() *CompletionContext
}

func (s *CompletionContext) CompletionContextCaster() *CompletionContext {
	_ = "STUB: not implemented"
	return nil
}
