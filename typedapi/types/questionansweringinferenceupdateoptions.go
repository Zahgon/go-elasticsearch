package types

type QuestionAnsweringInferenceUpdateOptions struct {
	MaxAnswerLength *int `json:"max_answer_length,omitempty"`

	NumTopClasses *int `json:"num_top_classes,omitempty"`

	Question string `json:"question"`

	ResultsField *string `json:"results_field,omitempty"`

	Tokenization *NlpTokenizationUpdateOptions `json:"tokenization,omitempty"`
}

func (s *QuestionAnsweringInferenceUpdateOptions) UnmarshalJSON(data []byte) error {
	_ = "STUB: not implemented"
	return nil
}

func NewQuestionAnsweringInferenceUpdateOptions() *QuestionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}

type QuestionAnsweringInferenceUpdateOptionsVariant interface {
	QuestionAnsweringInferenceUpdateOptionsCaster() *QuestionAnsweringInferenceUpdateOptions
}

func (s *QuestionAnsweringInferenceUpdateOptions) QuestionAnsweringInferenceUpdateOptionsCaster() *QuestionAnsweringInferenceUpdateOptions {
	_ = "STUB: not implemented"
	return nil
}
