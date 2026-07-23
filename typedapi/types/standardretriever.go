package types

type StandardRetriever struct {
	Collapse *FieldCollapse `json:"collapse,omitempty"`

	Filter []Query `json:"filter,omitempty"`

	MinScore *float32 `json:"min_score,omitempty"`

	Name_ *string `json:"_name,omitempty"`

	Query *Query `json:"query,omitempty"`

	SearchAfter []FieldValue `json:"search_after,omitempty"`

	Sort []SortCombinations `json:"sort,omitempty"`

	TerminateAfter *int `json:"terminate_after,omitempty"`
}

func (s *StandardRetriever) UnmarshalJSON(data []byte) error { _ = "STUB: not implemented"; return nil }

func NewStandardRetriever() *StandardRetriever { _ = "STUB: not implemented"; return nil }

type StandardRetrieverVariant interface {
	StandardRetrieverCaster() *StandardRetriever
}

func (s *StandardRetriever) StandardRetrieverCaster() *StandardRetriever {
	_ = "STUB: not implemented"
	return nil
}
