package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigCreateContainer struct {
	v *types.InferenceConfigCreateContainer
}

func NewInferenceConfigCreateContainer() *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) FillMask(fillmask types.FillMaskInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) LearningToRank(learningtorank types.LearningToRankConfigVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) Ner(ner types.NerInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) PassThrough(passthrough types.PassThroughInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) QuestionAnswering(questionanswering types.QuestionAnsweringInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) TextClassification(textclassification types.TextClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) TextEmbedding(textembedding types.TextEmbeddingInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) TextExpansion(textexpansion types.TextExpansionInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) ZeroShotClassification(zeroshotclassification types.ZeroShotClassificationInferenceOptionsVariant) *_inferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigCreateContainer) InferenceConfigCreateContainerCaster() *types.InferenceConfigCreateContainer {
	_ = "STUB: not implemented"
	return nil
}
