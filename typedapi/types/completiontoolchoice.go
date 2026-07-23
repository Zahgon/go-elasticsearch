package types

type CompletionToolChoice struct {
	Function CompletionToolChoiceFunction `json:"function"`

	Type string `json:"type"`
}

func (s *CompletionToolChoice) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompletionToolChoice() *CompletionToolChoice { _ = "STUB: not implemented"; return nil }

type CompletionToolChoiceVariant interface {
	CompletionToolChoiceCaster() *CompletionToolChoice
}

func (s *CompletionToolChoice) CompletionToolChoiceCaster() *CompletionToolChoice {
	_ = "STUB: not implemented"
	return nil
}

func (s *CompletionToolChoice) CompletionToolTypeCaster() *CompletionToolType {
	_ = "STUB: not implemented"
	return nil
}
