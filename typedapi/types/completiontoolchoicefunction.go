package types

type CompletionToolChoiceFunction struct {
	Name string `json:"name"`
}

func (s *CompletionToolChoiceFunction) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewCompletionToolChoiceFunction() *CompletionToolChoiceFunction {
	_ = "STUB: not implemented"
	return nil
}

type CompletionToolChoiceFunctionVariant interface {
	CompletionToolChoiceFunctionCaster() *CompletionToolChoiceFunction
}

func (s *CompletionToolChoiceFunction) CompletionToolChoiceFunctionCaster() *CompletionToolChoiceFunction {
	_ = "STUB: not implemented"
	return nil
}
