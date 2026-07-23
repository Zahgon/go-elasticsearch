package types

type QuestionAnsweringInferenceOptions struct {
	MaxAnswerLength *int `json:"max_answer_length,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *TokenizationConfigContainer `json:"tokenization,omitempty"`
}

func (s *QuestionAnsweringInferenceOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewQuestionAnsweringInferenceOptions() *QuestionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}

type QuestionAnsweringInferenceOptionsVariant interface {
	QuestionAnsweringInferenceOptionsCaster() *QuestionAnsweringInferenceOptions
}

func (s *QuestionAnsweringInferenceOptions) QuestionAnsweringInferenceOptionsCaster() *QuestionAnsweringInferenceOptions {
	_ = "STUB: not implemented"
	return nil
}
