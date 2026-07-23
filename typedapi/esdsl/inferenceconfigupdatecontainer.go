package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _inferenceConfigUpdateContainer struct {
	v *types.InferenceConfigUpdateContainer
}

func NewInferenceConfigUpdateContainer() *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) Classification(classification types.ClassificationInferenceOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) FillMask(fillmask types.FillMaskInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) Ner(ner types.NerInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) PassThrough(passthrough types.PassThroughInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) QuestionAnswering(questionanswering types.QuestionAnsweringInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) Regression(regression types.RegressionInferenceOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) TextClassification(textclassification types.TextClassificationInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) TextEmbedding(textembedding types.TextEmbeddingInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) TextExpansion(textexpansion types.TextExpansionInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) ZeroShotClassification(zeroshotclassification types.ZeroShotClassificationInferenceUpdateOptionsVariant) *_inferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}

func (s *_inferenceConfigUpdateContainer) InferenceConfigUpdateContainerCaster() *types.InferenceConfigUpdateContainer {
	_ = "STUB: not implemented"
	return nil
}
