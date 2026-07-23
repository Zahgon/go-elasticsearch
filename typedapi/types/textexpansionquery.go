package types

type TextExpansionQuery struct {
	Boost *float32 `json:"boost,omitempty"`

	ModelId string `json:"model_id"`

	ModelText string `json:"model_text"`

	PruningConfig *TokenPruningConfig `json:"pruning_config,omitempty"`
	QueryName_    *string             `json:"_name,omitempty"`
}

func (s *TextExpansionQuery) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewTextExpansionQuery() *TextExpansionQuery { _ = "STUB: not implemented"; return nil }

type TextExpansionQueryVariant interface {
	TextExpansionQueryCaster() *TextExpansionQuery
}

func (s *TextExpansionQuery) TextExpansionQueryCaster() *TextExpansionQuery {
	_ = "STUB: not implemented"
	return nil
}
