package esdsl

import "github.com/elastic/go-elasticsearch/v9/typedapi/types"

type _trainedModelTree struct {
	v *types.TrainedModelTree
}

func NewTrainedModelTree() *_trainedModelTree { _ = "STUB: not implemented"; return nil }

func (s *_trainedModelTree) ClassificationLabels(classificationlabels ...string) *_trainedModelTree {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTree) FeatureNames(featurenames ...string) *_trainedModelTree {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTree) TargetType(targettype string) *_trainedModelTree {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTree) TreeStructure(treestructures ...types.TrainedModelTreeNodeVariant) *_trainedModelTree {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTree) TreeStructureValues(treestructurevalues []types.TrainedModelTreeNode) *_trainedModelTree {
	_ = "STUB: not implemented"
	return nil
}

func (s *_trainedModelTree) TrainedModelTreeCaster() *types.TrainedModelTree {
	_ = "STUB: not implemented"
	return nil
}
