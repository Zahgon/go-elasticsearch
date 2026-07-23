package types

type InferenceConfigCreateContainer struct {
	Classification *ClassificationInferenceOptions `json:"classification,omitempty"`

	FillMask       *FillMaskInferenceOptions `json:"fill_mask,omitempty"`
	LearningToRank *LearningToRankConfig     `json:"learning_to_rank,omitempty"`

	Ner *NerInferenceOptions `json:"ner,omitempty"`

	PassThrough *PassThroughInferenceOptions `json:"pass_through,omitempty"`

	QuestionAnswering *QuestionAnsweringInferenceOptions `json:"question_answering,omitempty"`

	Regression *RegressionInferenceOptions `json:"regression,omitempty"`

	TextClassification *TextClassificationInferenceOptions `json:"text_classification,omitempty"`

	TextEmbedding *TextEmbeddingInferenceOptions `json:"text_embedding,omitempty"`

	TextExpansion *TextExpansionInferenceOptions `json:"text_expansion,omitempty"`

	ZeroShotClassification *ZeroShotClassificationInferenceOptions `json:"zero_shot_classification,omitempty"`
}

func NewInferenceConfigCreateContainer() *InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

type InferenceConfigCreateContainerVariant interface {
	InferenceConfigCreateContainerCaster() *InferenceConfigCreateContainer
}

func (s *InferenceConfigCreateContainer) InferenceConfigCreateContainerCaster() *InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}
