package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trainedModel struct {
	v *types.TrainedModel
}

func NewTrainedModel() *_trainedModel { _ = "STUB: not implemented"; return nil }

func (s *_trainedModel) Ensemble(ensemble types.EnsembleVariant) *_trainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModel) Tree(tree types.TrainedModelTreeVariant) *_trainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModel) TreeNode(treenode types.TrainedModelTreeNodeVariant) *_trainedModel {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModel) TrainedModelCaster() *types.TrainedModel {
	_ = "STUB: not implemented"
	return nil
}
