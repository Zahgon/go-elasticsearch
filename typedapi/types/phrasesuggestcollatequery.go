package types

type PhraseSuggestCollateQuery struct {
	Id *string `json:"id,omitempty"`

	Source ScriptSource `json:"source,omitempty"`
}

func (s *PhraseSuggestCollateQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewPhraseSuggestCollateQuery() *PhraseSuggestCollateQuery {
	_ = "STUB: not implemented"
	return nil
}

type PhraseSuggestCollateQueryVariant interface {
	PhraseSuggestCollateQueryCaster() *PhraseSuggestCollateQuery
}

func (s *PhraseSuggestCollateQuery) PhraseSuggestCollateQueryCaster() *PhraseSuggestCollateQuery {
	_ = "STUB: not implemented"
	return nil
}
