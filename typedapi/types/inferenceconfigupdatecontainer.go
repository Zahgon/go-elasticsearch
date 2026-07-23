package types

type InferenceConfigUpdateContainer struct {
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`

	FillMask *FillMaskInferenceUpdateOptions `json:"fill_mask,omitempty"`

	Ner *NerInferenceUpdateOptions `json:"ner,omitempty"`

	PassThrough *PassThroughInferenceUpdateOptions `json:"pass_through,omitempty"`

	QuestionAnswering *QuestionAnsweringInferenceUpdateOptions `json:"question_answering,omitempty"`

	Regression *RegressionInferenceOptions `json:"regression,omitempty"`

	TextClassification *TextClassificationInferenceUpdateOptions `json:"text_classification,omitempty"`

	TextEmbedding *TextEmbeddingInferenceUpdateOptions `json:"text_embedding,omitempty"`

	TextExpansion *TextExpansionInferenceUpdateOptions `json:"text_expansion,omitempty"`

	ZeroShotClassification *ZeroShotClassificationInferenceUpdateOptions `json:"zero_shot_classification,omitempty"`
}

func NewInferenceConfigUpdateContainer() *InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

type InferenceConfigUpdateContainerVariant interface {
	InferenceConfigUpdateContainerCaster() *InferenceConfigUpdateContainer
}

func (s *InferenceConfigUpdateContainer) InferenceConfigUpdateContainerCaster() *InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}
